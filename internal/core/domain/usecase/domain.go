package domain

import "totem-pedidos/internal/core/ports"

type OrderUsecase struct {
	Repository ports.DatabaseRepository
}

func New() *OrderUsecase {
	return &OrderUsecase{}
}

func (o *OrderUsecase) CreateOrder() error {
	o.Repository.Get()
	return nil
}
