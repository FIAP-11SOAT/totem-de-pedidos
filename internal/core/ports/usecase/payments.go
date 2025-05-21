package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type Payments interface {
	GetPaymentByID(paymentID string) (*entity.Payment, error)
	PaymentWebHook(*mercadopago.WebhookPayload) error
}
