package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
)

type Order interface {
	CreateOrder(input entity.Order) (int, error)
	UpdateOrderStatus(id int, status string) error
	GetOrderByID(id int) (entity.Order, error)
	ListOrders(filter input.OrderFilterInput) ([]entity.Order, error)
	Checkout(orderID int) error
}
