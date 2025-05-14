package usecase

import (
	"context"
	"fmt"
	"time"

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

func (p *Product) CreateProduct(productInput *usecase.ProductInput) (*entity.Product, error) {
	category, err := p.getCategoryName(productInput.CategoryName)
	if err != nil {
		return nil, err
	}

	productToCreate := &entity.Product{
		Name:            productInput.Name,
		Description:     productInput.Description,
		Price:           productInput.Price,
		ImageURL:        productInput.ImageURL,
		PreparationTime: 0,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		CategoryID:      category.ID,
	}

	createdProductId, err := p.productRepository.CreateProduct(context.Background(), productToCreate)
	if err != nil {
		return nil, fmt.Errorf("error creating product")
	}

	return &entity.Product{
		ID:          createdProductId,
		Description: productInput.Description,
		Price:       productInput.Price,
	}, nil
}

func (p *Product) getCategoryName(categoryName string) (*entity.ProductCategory, error) {
	category, err := p.productRepository.GetCategoryByName(categoryName)
	if err != nil {
		return nil, fmt.Errorf("category not found")
	}

	return category, nil
}

func (p *Product) GetProductByCategoryID(categoryID int) ([]*entity.Product, error) {
	products, err := p.productRepository.GetProductsByCategoryID(context.Background(), categoryID)
	if err != nil {
		return nil, fmt.Errorf("error fetching products by category ID: %w", err)
	}

	return products, nil
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
