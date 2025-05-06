package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
)

type Customer struct {
	CustomerRepository repositories.Customer
}

// NewCustomerUseCase cria uma nova instância do caso de uso de cliente com o repositório fornecido.
func NewCustomerUseCase(repository repositories.Customer) usecase.Customer {
	return &Customer{CustomerRepository: repository}
}

func (p *Customer) CreateCustomer(customerDTO *usecase.CustomerInput) (*entity.Customer, error) {
	customerEntity := entity.Customer{
		ID:    customerDTO.ID,
		Name:  customerDTO.Nome,
		Email: customerDTO.Email,
		TaxID: customerDTO.TaxID,
	}

	p.CustomerRepository.CreateCustomer(&customerEntity)
	return nil, nil
}
