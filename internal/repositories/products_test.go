package repositories_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/pkg/tests"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	connStr := tests.CreatePostgresDataBase(t)

	// TODO: fix
	// dbConnection := dbadapter.New(
	// 	dbadapter.Input{
	// 		Db_driver:  "postgres",
	// 		Db_user:    "totempedidos",
	// 		Db_pass:    "totempedidos",
	// 		Db_host:    "totempedidos",
	// 		Db_name:    "totempedidos",
	// 		Db_options: "?sslmode=disable",
	// 	},
	// )

	client, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		panic(err)
	}

	productRepository := repositories.NewProductRepository(&dbadapter.DatabaseAdapter{Client: client})

	t.Run("should failed with empty Product entity", func(t *testing.T) {
		_, err := productRepository.CreateProduct(context.Background(), &entity.Product{})
		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}
