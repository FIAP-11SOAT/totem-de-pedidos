package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type ProductInput struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageURL     string  `json:"imageUrl"`
	CategoryName string  `json:"categoryName"`
}

type Product interface {
	GetProductById(id string) (*entity.Product, error)
	GetProducts(description string) ([]*entity.Product, error)
	CreateProduct(productDTO *ProductInput) (*entity.Product, error)
	UpdateProduct(productDTO *ProductInput) (*entity.Product, error)
	DeleteProduct(productID string) error
	GetCategories() ([]*entity.ProductCategory, error)
	GetProductByCategoryID(categoryID int) ([]*entity.Product, error)
}

func (p *ProductInput) Validate() error {
	// TODO: implement-me
	return nil
}
