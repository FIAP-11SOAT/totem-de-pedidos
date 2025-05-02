package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type ProductInput struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IDCategory  string  `json:"category_id"`
	Category    string  `json:"category"`
}

type Product interface {
	GetProductById(id string) (*entity.Product, error)
	GetProducts(description string) ([]*entity.Product, error)
	CreateProduct(productDTO *ProductInput) (*entity.Product, error)
	UpdateProduct(productDTO *ProductInput) (*entity.Product, error)
	DeleteProduct(productID string) error
	GetCategoryName(categoryName string) (string, error)
	GetCategories() ([]*entity.ProductCategory, error)
}
