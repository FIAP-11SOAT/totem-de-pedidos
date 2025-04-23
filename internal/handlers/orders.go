package handlers

import (
	"net/http"
	"totem-pedidos/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Usecase ports.UsecaseInterface
}

func NewOrderRouter() *OrderHandler {
	return &OrderHandler{}
}

func (o *OrderHandler) CreateOrder(c echo.Context) error {

	if err := o.Usecase.CreateOrder(); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating order")
	}

	return c.String(http.StatusOK, "Hello, World!")
}
