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

func TestCreateCategory(t *testing.T) {
	connStr := tests.CreatePostgresDataBase(t)

	client, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}

	categoryRepo := repositories.NewCategoryRepository(&dbadapter.DatabaseAdapter{Client: client})

	t.Run("should fail with empty category entity", func(t *testing.T) {
		_, err := categoryRepo.CreateCategory(context.Background(), entity.ProductCategory{})
		assert.Error(t, err)
		assert.NotNil(t, err)
	})

	t.Run("should create category successfully", func(t *testing.T) {
		now := time.Now().UTC()

		category := entity.ProductCategory{
			Name:        "Lanches",
			Description: "Categoria de Lanches",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		categoryId, err := categoryRepo.CreateCategory(context.Background(), category)
		assert.NoError(t, err)
		assert.True(t, categoryId > 0)
	})
}
