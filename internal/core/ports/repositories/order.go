package repositories

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
)

type Order interface {
	CreateOrder(entity.Order) (int, error)
	UpdateStatus(orderID int, status string) error
	GetOrderByID(orderID int) (entity.Order, error)
	ListOrders(filter input.OrderFilterInput) ([]entity.Order, error)
}
