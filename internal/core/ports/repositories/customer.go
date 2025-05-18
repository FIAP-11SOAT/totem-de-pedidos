package repositories

import "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"

type Customer interface {
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	IdentifyCustomer(taxID *string) (*entity.Customer, error)
}
