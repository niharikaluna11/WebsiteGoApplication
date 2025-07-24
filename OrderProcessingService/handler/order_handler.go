package handlers

import (
	"OrderProcessingService/models"
	pubsubevents "OrderProcessingService/pubsub"
	"OrderProcessingService/services"
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/kataras/iris/v12"
)

func NewOrderHandler(service services.OrderServiceInterface, pubsubClient *pubsub.Client, ctx context.Context) *OrderHandler {
	return &OrderHandler{
		Service:      service,
		PubSubClient: pubsubClient,
		Ctx:          ctx,
	}
}

type OrderHandler struct {
	Service      services.OrderServiceInterface
	PubSubClient *pubsub.Client
	Ctx          context.Context
}

func (h *OrderHandler) CreateOrder(ctx iris.Context) {

	var dto models.OrderCreateDTO
	if err := ctx.ReadJSON(&dto); err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		return
	}

	order, err := h.Service.GenerateOrder(&dto)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	event := &models.OrderEvent{
		OrderID:     order.OrderId,
		TotalAmount: order.TotalAmount,
	}

	err = pubsubevents.PublishOrderCreated(h.Ctx, h.PubSubClient, *event)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to publish order event", "details": err.Error()})
		return
	}
	print("Order event published successfully\n")

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Order created successfully", "result": order})
}

func (h *OrderHandler) GetOrderById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	order, err := h.Service.GetOrderById(id)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}
	if order == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Order not found"})
		return
	}
	ctx.JSON(order)
}

func (h *OrderHandler) UpdateOrderStatus(ctx iris.Context) {
	var dto models.OrderStatusUpdateDTO
	if err := ctx.ReadJSON(&dto); err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		return
	}
	order, err := h.Service.UpdateOrderStatus(dto.OrderId, string(dto.Status))
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(order)
}
