package repositories

import "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"

type Product interface {
	ListProducts(description string) ([]*entity.Product, error)
	FindProductById(id string) (*entity.Product, error)
	CreateProduct(product *entity.Product) (*entity.Product, error)
	GetCategoryDescription(categoryName string) (string, error)
	GetCategories() ([]*entity.ProductCategory, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(productID string) error
}
