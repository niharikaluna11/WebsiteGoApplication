package main

import (
	"OrderProcessingService/config"
	handlers "OrderProcessingService/handler"
	"OrderProcessingService/middleware"
	repository "OrderProcessingService/respository"
	"OrderProcessingService/services"
	"OrderProcessingService/subevents"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found ⚠️")
	}
}

func main() {
	log.Println("Starting Order Processing Service...")

	app := iris.New()
	db, error := config.ConnectToDB()
	if error != nil {
		fmt.Printf("Connection Lost: %s", error)
		return
	}

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "niharikaprojects")
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err)
	}

	orderRepo := repository.NewOrderRepoImpl(db)
	orderService := &services.OrderService{Repo: orderRepo}
	orderHandler := handlers.NewOrderHandler(orderService, pubsubClient, ctx)

	userRepo := repository.NewUserRepoImpl(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Welcome to the Order Processing Service!")
	})

	repo := &repository.OrderRepoImpl{DB: db}

	go subevents.ListenForPayments(repo)

	userAPI := app.Party("/users")
	{
		userAPI.Post("/", userHandler.Register)
		userAPI.Post("/login", userHandler.Login)
	}

	orderAPI := app.Party("/orders", middleware.JWTMiddleware)
	{
		orderAPI.Post("/", orderHandler.CreateOrder)
		orderAPI.Get("/{id:string}", orderHandler.GetOrderById)
		orderAPI.Patch("/{id:string}/status", orderHandler.UpdateOrderStatus)
	}

	log.Println("Connected to the database successfully.")

	app.Listen(":8081")
}
