package handlers

import (
	"net/http"
	"totem-pedidos/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Repository ports.DatabaseRepository
	Usecase    ports.UsecaseInterface
}

func NewOrderRouter(repository ports.DatabaseRepository, usecase ports.UsecaseInterface) *OrderHandler {
	return &OrderHandler{Repository: repository, Usecase: usecase}
}

func (o *OrderHandler) CreateOrder(c echo.Context) error {

	// DTO create order
	// bind
	// valida
 
	o.Usecase.CreateOrder()

	return c.String(http.StatusOK, "Hello, World!")
}
