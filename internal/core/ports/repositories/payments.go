package repositories

import "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"

type Payments interface {
	GetPaymentByID(paymentID string) (*entity.Payment, error)
}
