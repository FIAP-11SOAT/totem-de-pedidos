package service

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/services/payment"
)

type PaymentService interface {
	Payment(input payment.PaymentInput) (payment.PaymentOutput, error)
}
