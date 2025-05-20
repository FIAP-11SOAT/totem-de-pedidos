package usecase

import "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"

type Payments interface {
	GetPaymentByID(paymentID string) (string, error)
	PaymentWebHook(*mercadopago.WebhookPayload) error
}
