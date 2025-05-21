package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/repositories"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
)

type Payments struct {
	paymentRepository repositories.Payments
}

func NewPaymentsUseCase(paymentRepository repositories.Payments) usecase.Payments {
	return &Payments{
		paymentRepository: paymentRepository,
	}
}

func (p *Payments) GetPaymentByID(paymentID string) (*entity.Payment, error) {
	payment, err := p.paymentRepository.GetPaymentByID(paymentID)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (p *Payments) PaymentWebHook(payload *mercadopago.WebhookPayload) error {
	return nil
}
