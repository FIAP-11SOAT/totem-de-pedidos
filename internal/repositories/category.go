package repositories

import (
	"context"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type categoryRepository struct {
	sqlClient *pgx.Conn
}

func NewCategoryRepository(database *dbadapter.DatabaseAdapter) repositories.Category {
	return &categoryRepository{
		sqlClient: database.Client,
	}
}

func (c *categoryRepository) CreateCategory(ctx context.Context, category entity.ProductCategory) (int, error) {
	var createdId int

	err := c.sqlClient.QueryRow(ctx, createCategoryQuery(),
		category.Name,
		category.CreatedAt,
		category.UpdatedAt,
	).Scan(&createdId)

	if err != nil {
		return 0, err
	}

	return createdId, nil
}

func createCategoryQuery() string {
	return `
		INSERT INTO product_categories (
			name,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}
