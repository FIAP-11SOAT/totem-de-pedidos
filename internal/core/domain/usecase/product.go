package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"

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
	product, err := p.productRepository.FindProductById(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("error fetching product by id: %w", err)
	}
	return product, nil
}

func (p *Product) GetProducts(input *input.ProductFilterInput) ([]*entity.Product, error) {
	products, err := p.productRepository.ListProducts(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) CreateProduct(productInput *input.ProductInput) (*entity.Product, error) {
	productToCreate := &entity.Product{
		Name:            productInput.Name,
		Description:     productInput.Description,
		Price:           productInput.Price,
		ImageURL:        productInput.ImageURL,
		PreparationTime: 0, // ajuste se necessário
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		CategoryID:      productInput.CategoryID,
	}

	createdProductId, err := p.productRepository.CreateProduct(context.Background(), productToCreate)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	// Retorna o produto completo
	productToCreate.ID = createdProductId
	return productToCreate, nil
}

func (p *Product) getCategoryName(categoryName string) (*entity.ProductCategory, error) {
	category, err := p.productRepository.GetCategoryByName(context.Background(), categoryName)
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

func (p *Product) UpdateProduct(productDTO *input.ProductInput) (*entity.Product, error) {
	return nil, nil
}

func (p *Product) DeleteProduct(productID string) error {
	return nil
}

func (p *Product) GetCategories() ([]*entity.ProductCategory, error) {
	return nil, nil
}
