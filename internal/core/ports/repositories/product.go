package repositories

import (
	"context"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type Product interface {
	CreateProduct(ctx context.Context, product entity.Product) (int, error)
	DeleteProduct(productID string) error
	FindProductById(id string) (*entity.Product, error)
	GetCategories() ([]entity.ProductCategory, error)
	GetCategoryByName(categoryName string) (entity.ProductCategory, error)
	ListProducts(description string) ([]*entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
}
