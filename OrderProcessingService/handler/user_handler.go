package handlers

import (
	"OrderProcessingService/models"
	"OrderProcessingService/services"
	"net/http"

	"github.com/kataras/iris/v12"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) Register(ctx iris.Context) {
	var dto models.UserRegisterDTO
	if err := ctx.ReadJSON(&dto); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if err := h.svc.Register(&dto); err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(iris.Map{"message": "user registered successfully"})
}

func (h *UserHandler) Login(ctx iris.Context) {
	var dto models.UserLoginDTO
	if err := ctx.ReadJSON(&dto); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	token, err := h.svc.Login(&dto)
	if err != nil {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"token": token, "message": "login successful"})
}
