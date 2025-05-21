package mock

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
)

type OrderMock struct {
	CreateOrderFunc  func(entity.Order) (int, error)
	UpdateStatusFunc func(int, string) error
	GetOrderByIDFunc func(int) (entity.Order, error)
	ListOrdersFunc   func(input.OrderFilterInput) ([]entity.Order, error)
}

func NewOrderRepositoryMock() *OrderMock {
	return &OrderMock{
		CreateOrderFunc:  func(entity.Order) (int, error) { return 0, nil },
		UpdateStatusFunc: func(int, string) error { return nil },
		GetOrderByIDFunc: func(i int) (entity.Order, error) { return entity.Order{}, nil },
		ListOrdersFunc:   func(ofi input.OrderFilterInput) ([]entity.Order, error) { return []entity.Order{}, nil },
	}
}

func (o *OrderMock) CreateOrder(order entity.Order) (int, error) { return o.CreateOrderFunc(order) }

func (o *OrderMock) UpdateStatus(i int, s string) error { return o.UpdateStatusFunc(i, s) }

func (o *OrderMock) GetOrderByID(i int) (entity.Order, error) { return o.GetOrderByIDFunc(i) }

func (o *OrderMock) ListOrders(ofi input.OrderFilterInput) ([]entity.Order, error) {
	return o.ListOrdersFunc(ofi)
}
