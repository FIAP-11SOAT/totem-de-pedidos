package repositories

import (
	"context"
	"fmt"

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

func (s *CustomerRepository) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	query := `INSERT INTO customers (id, name, email, tax_id) VALUES (@id, @name, @email, @taxid)`
	args := pgx.NamedArgs{
		"id":    customer.ID,
		"name":  customer.Name,
		"email": customer.Email,
		"taxid": customer.TaxID,
	}

	_, err := s.sqlClient.Exec(context.Background(), query, args)
	if err != nil {
		return nil, fmt.Errorf("error on inserting Customer: %w", err)
	}

	return customer, nil
}
