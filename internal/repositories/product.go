package repositories

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type productRepository struct {
	sqlClient *pgx.Conn
}

// NewProductRepository cria uma nova instância de productRepository utilizando o adaptador de banco de dados fornecido.
func NewProductRepository(database *dbadapter.DatabaseAdapter) repositories.Product {
	return &productRepository{
		sqlClient: database.Client(),
	}
}

func (s *productRepository) ListProducts(description string) ([]*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (s *productRepository) FindProductById(id string) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (s *productRepository) CreateProduct(product *entity.Product) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (s *productRepository) GetCategoryDescription(categoryDescription string) (string, error) {
	// TODO: implement-me
	return "", nil
}

func (s *productRepository) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	// TODO: implement-me
	return nil, nil
}

func (s *productRepository) DeleteProduct(productID string) error {
	// TODO: implement-me
	return nil
}

// GetCategories returns a list of all product categories
func (s *productRepository) GetCategories() ([]*entity.ProductCategory, error) {
	// TODO: implement-me
	return nil, nil
}
