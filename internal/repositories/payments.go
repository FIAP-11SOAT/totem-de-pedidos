package repositories

import (
	"context"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/jackc/pgx/v5"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
)

type paymentsRepository struct {
	sqlClient *pgx.Conn
}

func NewPaymentsRepository(database *dbadapter.DatabaseAdapter) repositories.Payments {
	return &paymentsRepository{
		sqlClient: database.Client,
	}
}

func (r *paymentsRepository) GetPaymentByID(paymentID string) (*entity.Payment, error) {
	query := `SELECT id, amount, payment_date, status, provider FROM payments WHERE id = $1`
	var payment entity.Payment
	err := r.sqlClient.
		QueryRow(context.Background(), query, paymentID).
		Scan(
			&payment.ID,
			&payment.Amount,
			&payment.PaymentDate,
			&payment.Status,
			&payment.Provider,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &payment, nil
		}
		return &payment, err
	}
	return &payment, nil
}
