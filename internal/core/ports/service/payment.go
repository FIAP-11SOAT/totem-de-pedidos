package service

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/output"
)

type PaymentService interface {
	GeneratePaymentQrCode(input *input.CreatePaymentInput) (*output.CreatePaymentOutput, error)
	GetPaymentByID(paymentID string) (*output.GetPaymentOutput, error)
}
