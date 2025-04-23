package api

import (
	"totem-pedidos/internal/core/ports"
	"totem-pedidos/internal/handlers"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Repository ports.DatabaseRepository
	Usecase    ports.UsecaseInterface
}

func NewOrderHandler(repository ports.DatabaseRepository, usecase ports.UsecaseInterface) *OrderHandler {
	return &OrderHandler{Repository: repository, Usecase: usecase}
}

func (o *OrderHandler) OrderRouters(e *echo.Echo) {
	e.GET("/", handlers.CreateOrder)
	e.GET("/", handlers.CreateOrder)
}
