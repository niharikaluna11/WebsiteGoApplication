package handlers

import (
	"PaymentProcessingService/models"
	"PaymentProcessingService/services"
	"fmt"

	"github.com/kataras/iris/v12"
)

type PaymentHandler struct {
	Service services.PaymentService
}

func (ph *PaymentHandler) GetPaymentById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	result, err := ph.Service.GetPaymentById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
		return
	}

	if result.OrderID == "" {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Payment not found with PAYMENT ID: %s", id)})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Payment retrieved successfully", "payment": result})
}

func (ph *PaymentHandler) UpdatePaymentStatus(ctx iris.Context) {
	var paymentStatusUpdateDTO models.PaymentStatusUpdateDTO
	if err := ctx.ReadJSON(&paymentStatusUpdateDTO); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	result, err := ph.Service.UpdatePaymentStatus(paymentStatusUpdateDTO.OrderID, string(paymentStatusUpdateDTO.Status))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
		return
	}

	if result == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "Payment not found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Payment status updated successfully to " + result.Status})
}
