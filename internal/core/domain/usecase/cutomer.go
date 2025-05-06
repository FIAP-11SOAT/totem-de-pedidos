package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
)

type Customer struct {
	CustomerRepository repositories.Customer
}

func NewCustomerUseCase(repository repositories.Customer) usecase.Customer {
	return &Customer{CustomerRepository: repository}
}

func (p *Customer) CreateCustomer(CustomerDTO *usecase.CustomerInput) (*entity.Customer, error) {
	return nil, nil
}
