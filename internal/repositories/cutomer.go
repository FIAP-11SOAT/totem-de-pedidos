package repositories

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type CustomerRepository struct {
	sqlClient *pgx.Conn
}

func NewCustomerRepository(database *dbadapter.DatabaseAdapter) repositories.Customer {
	return &CustomerRepository{
		sqlClient: database.Client(),
	}
}

func (s *CustomerRepository) CreateCustomer(Customer *entity.Customer) (*entity.Customer, error) {
	return nil, nil
}
