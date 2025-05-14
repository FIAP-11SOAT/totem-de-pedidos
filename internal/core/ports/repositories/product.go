package repositories

import (
	"context"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type Product interface {
	ListProducts(description string) ([]*entity.Product, error)
	FindProductById(id string) (*entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.Product) (int, error)
	GetCategoryByName(categoryName string) (*entity.ProductCategory, error)
	GetCategories() ([]*entity.ProductCategory, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(productID string) error
	GetProductsByCategoryID(ctx context.Context, categoryID int) ([]*entity.Product, error)
}
