package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/inputs"
)

type Product interface {
	GetProductById(id string) (*entity.Product, error)
	GetProducts(*inputs.ProductFilterInput) ([]*entity.Product, error)
	CreateProduct(productDTO *inputs.ProductInput) (*entity.Product, error)
	UpdateProduct(productDTO *inputs.ProductInput) (*entity.Product, error)
	DeleteProduct(productID string) error

	GetCategories() ([]*entity.ProductCategory, error)
	GetProductByCategoryID(categoryID int) ([]*entity.Product, error)
}
