package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type CustomerInput struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Email string  `json:"email"`
	TaxID float64 `json:"taxid"`
}

type Customer interface {
	CreateCustomer(customerDTO *CustomerInput) (*entity.Customer, error)
}
