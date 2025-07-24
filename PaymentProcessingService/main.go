package main

import (
	"PaymentProcessingService/config"
	handlers "PaymentProcessingService/handler"
	"PaymentProcessingService/repository"
	"PaymentProcessingService/services"
	"PaymentProcessingService/subevents"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}
}

func main() {
	app := iris.New()

	db, err := config.ConnectToDB()
	if err != nil {
		fmt.Println("Connection Lost", err)
		return
	}

	repo := &repository.PaymentRepoImpl{DB: db}

	go subevents.SubscribeToOrderEvents(repo)

	service := &services.PaymentService{Repo: repo}
	paymentHandler := &handlers.PaymentHandler{Service: *service}

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Payment Service!")
	})

	app.Get("/api/v1/getpayment/:id", paymentHandler.GetPaymentById)

	fmt.Println("🚀 Payment Service running on http://localhost:8082")
	app.Listen(":8082")
}
