package repositories

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type paymentsRepository struct {
	sqlClient *pgx.Conn
}

func NewPaymentsRepository(sqlClient *pgx.Conn) repositories.Payments {
	return &paymentsRepository{
		sqlClient: sqlClient,
	}
}

func (r *paymentsRepository) GetPaymentByID(paymentID string) (string, error) {
	query := `SELECT payment_id FROM payments WHERE payment_id = $1`
	return "", nil
}
