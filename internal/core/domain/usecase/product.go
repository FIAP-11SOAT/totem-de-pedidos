package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
)

type Product struct {
	productRepository repositories.Product
}

func NewProductUseCase(repository repositories.Product) usecase.Product {
	return &Product{productRepository: repository}
}

func (p *Product) GetProductById(id string) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *Product) GetProducts(description string) ([]*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *Product) CreateProduct(productDTO *usecase.ProductInput) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *Product) GetCategoryName(categoryName string) (string, error) {
	// TODO: implement-me
	return "", nil
}

func (p *Product) UpdateProduct(productDTO *usecase.ProductInput) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *Product) DeleteProduct(productID string) error {
	// TODO: implement-me
	return nil
}

func (p *Product) GetCategories() ([]*entity.ProductCategory, error) {
	// TODO: implement-me
	return nil, nil
}
