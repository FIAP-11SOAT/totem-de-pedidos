package mapper

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/output"
)

func MapPaymentToOutput(payment *entity.Payment) output.PaymentOutput {
	out := output.PaymentOutput{
		ID:          payment.ID,
		Amount:      payment.Amount,
		PaymentDate: payment.PaymentDate,
		Status:      string(payment.Status),
		Provider:    payment.Provider,
	}
	return out
}
