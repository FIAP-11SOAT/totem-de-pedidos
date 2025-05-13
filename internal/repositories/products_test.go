package repositories_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/pkg/tests"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	connStr := tests.CreatePostgresDataBase(t)

	client, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}

	productRepo := repositories.NewProductRepository(&dbadapter.DatabaseAdapter{Client: client})
	categoryRepo := repositories.NewCategoryRepository(&dbadapter.DatabaseAdapter{Client: client})

	t.Run("should fail with empty Product entity", func(t *testing.T) {
		_, err := productRepo.CreateProduct(context.Background(), entity.Product{})
		assert.Error(t, err)
		assert.NotNil(t, err)
	})

	t.Run("should create a product successfully", func(t *testing.T) {
		category := entity.ProductCategory{
			Name:        "test",
			Description: "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		categoryId, err := categoryRepo.CreateCategory(context.Background(), category)
		assert.NoError(t, err)

		product := entity.Product{
			Name:            "test",
			Description:     "test",
			Price:           1.0,
			ImageURL:        "https://example.com/image.png",
			PreparationTime: 0,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			CategoryID:      categoryId,
		}

		productId, err := productRepo.CreateProduct(context.Background(), product)
		assert.NoError(t, err)
		assert.True(t, productId > 0)

	})
}
