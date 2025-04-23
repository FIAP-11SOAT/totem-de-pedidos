package api

import (
	"totem-pedidos/internal/core/ports"
	"totem-pedidos/internal/handlers"

	"github.com/labstack/echo/v4"
)

type OrderRouter struct {
	Repository ports.DatabaseRepository
	Usecase    ports.UsecaseInterface
}

func NewOrderRouter(repository ports.DatabaseRepository, usecase ports.UsecaseInterface) *OrderRouter {
	return &OrderRouter{Repository: repository, Usecase: usecase}
}

func (o *OrderRouter) OrderRouters(e *echo.Echo) {
	orderHandler := handlers.NewOrderRouter(o.Repository, o.Usecase)

	e.GET("/", orderHandler.CreateOrder)
}
