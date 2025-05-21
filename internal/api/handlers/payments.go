package handlers

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/mapper"
	"github.com/labstack/echo/v4"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/helper"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
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
		return c.JSON(400, helper.HttpResponse{
			Message: "Payment ID is required",
		})
	}
	payment, err := h.paymentService.GetPaymentByID(paymentID)
	if err != nil {
		c.Logger().Error(helper.LoggerInfo{
			Scope:   "get:payment:id",
			Message: "Failed to get payment by ID",
		})
		return c.JSON(500, helper.HttpResponse{
			Message: "Internal server error",
		})
	}
	return c.JSON(200, mapper.MapPaymentToOutput(payment))
}

func (h *PaymentHandler) PaymentWebHook(c echo.Context) error {
	//isValid := mercadopago.CheckWebhookSignature(mercadopago.CheckWebhookSignatureInput{
	//	XSignature: c.Request().Header.Get("X-Signature"),
	//	XRequestId: c.Request().Header.Get("X-Request-Id"),
	//	DataID:     c.QueryParam("data.id"),
	//	Secret:     os.Getenv("MP_WEBHOOK_SECRET"),
	//})
	//
	//if !isValid {
	//	c.Logger().Error(helper.LoggerInfo{
	//		Scope:   "post:payment:webhook",
	//		Message: "Invalid webhook signature",
	//	})
	//	return c.NoContent(401)
	//}

	c.Logger().Info(helper.LoggerInfo{
		Scope:   "post:payment:webhook",
		Message: "Webhook payload received",
	})
	var content map[string]interface{}
	if err := c.Bind(&content); err != nil {
		c.Logger().Error(helper.LoggerInfo{
			Scope:   "post:payment:webhook",
			Message: "Failed to bind webhook content",
		})
		return c.NoContent(400)
	}

	payload := new(mercadopago.WebhookPayload)
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(helper.LoggerInfo{
			Scope:   "post:payment:webhook",
			Message: "Failed to bind webhook payload",
		})
		return c.NoContent(400)
	}

	err := h.paymentService.PaymentWebHook(payload)
	if err != nil {
		c.Logger().Error(helper.LoggerInfo{
			Scope:   "post:payment:webhook",
			Message: "Failed to process webhook payload",
		})
		return c.NoContent(500)
	}

	c.Logger().Info(payload)
	c.Logger().Info(helper.LoggerInfo{
		Scope:   "post:payment:webhook",
		Message: "Webhook payload processed successfully",
	})
	return c.JSON(200, helper.HttpResponse{
		Message: "Webhook processed successfully",
	})
}
