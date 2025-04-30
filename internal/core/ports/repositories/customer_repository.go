package repositories

import (
	"totem-pedidos/internal/core/domain/entity"
)

type CustomerRepository interface {
	GetByTaxId(id int) (*entity.Customer, error)
	Create(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
}
