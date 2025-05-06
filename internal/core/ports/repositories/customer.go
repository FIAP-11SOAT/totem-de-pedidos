package repositories

import "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"

type Customer interface {
	CreateCustomer(product *entity.Customer) (*entity.Customer, error)
}
