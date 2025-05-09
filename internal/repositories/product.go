package repositories

import (
	"context"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type productRepository struct {
	sqlClient *pgx.Conn
}

func NewProductRepository(database *dbadapter.DatabaseAdapter) repositories.Product {
	return &productRepository{
		sqlClient: database.Client,
	}
}

func (p *productRepository) ListProducts(description string) ([]*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *productRepository) FindProductById(id string) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product *entity.Product) (int, error) {
	var createdProductId int

	err := p.sqlClient.QueryRow(ctx, createProductQuery(),
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.ImageURL,
		product.PreparationTime,
		product.CreatedAt,
		product.UpdatedAt,
		product.CategoryID,
	).Scan(&createdProductId)
	if err != nil {
		return 0, err
	}

	return createdProductId, nil
}

func createProductQuery() string {
	return `INSERT....RETURNING id`
}

func (p *productRepository) GetCategoryByName(ctx context.Context, categoryName string) (*entity.ProductCategory, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM product_categories
		WHERE name = $1
	`

	var category entity.ProductCategory
	err := p.sqlClient.QueryRow(ctx, query, categoryName).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	query := `
		UPDATE products
		SET name = $1,
			description = $2,
			price = $3,
			image_url = $4,
			preparation_time = $5,
			updated_at = $6,
			category_id = $7
		WHERE id = $8
		RETURNING id, name, description, price, image_url, preparation_time, created_at, updated_at, category_id
	`

	var updatedProduct entity.Product
	err := p.sqlClient.QueryRow(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.ImageURL,
		product.PreparationTime,
		product.UpdatedAt,
		product.CategoryID,
		product.ID,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.Description,
		&updatedProduct.Price,
		&updatedProduct.ImageURL,
		&updatedProduct.PreparationTime,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt,
		&updatedProduct.CategoryID,
	)

	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, productID string) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	
	_, err := p.sqlClient.Exec(ctx, query, productID)
	if err != nil {
		return err
	}

	return nil
}

// GetCategories returns a list of all product categories
func (p *productRepository) GetCategories() ([]*entity.ProductCategory, error) {
	// TODO: implement-me
	return nil, nil
}
