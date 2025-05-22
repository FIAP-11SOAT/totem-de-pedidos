package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/output"
)

type Payments interface {
	CreatePayment(input input.CreatePaymentInput) (*output.CreatePaymentOutput, error)
	GetPaymentByID(paymentID string) (*entity.Payment, error)
	PaymentWebHook(*mercadopago.WebhookPayload) error
}
