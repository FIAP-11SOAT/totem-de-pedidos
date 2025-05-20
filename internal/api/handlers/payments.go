package handlers

import (
	"os"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService usecase.Payments
}

func NewPaymentHandler(paymentService usecase.Payments) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) GetByID(c echo.Context) error {
	paymentID := c.Param("id")
	if paymentID == "" {
		return c.JSON(400, map[string]string{"error": "paymentID is required"})
	}

	// status, err := h.PaymentService.GetPaymentByID(paymentID)
	// if err != nil {
	// 	return c.JSON(500, map[string]string{"error": "failed to get payment status"})
	// }

	return c.JSON(200, map[string]string{"status": paymentID})
}

func (h *PaymentHandler) PaymentWebHook(c echo.Context) error {
	isValid := mercadopago.CheckWebhookSignature(mercadopago.CheckWebhookSignatureInput{
		XSignature: c.Request().Header.Get("X-Signature"),
		XRequestId: c.Request().Header.Get("X-Request-Id"),
		DataID:     c.QueryParam("data.id"),
		Secret:     os.Getenv("MP_WEBHOOK_SECRET"),
	})

	if !isValid {
		return c.JSON(400, map[string]string{"error": "invalid signature"})
	}

	payload := new(mercadopago.WebhookPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid payload"})
	}

	err := h.paymentService.PaymentWebHook(payload)
	if err != nil {
		c.Logger().Errorf("error processing webhook: %v", err)
	}

	return c.JSON(200, map[string]string{"message": "webhook received"})
}
