package payment

type PaymentInput struct {
	Amount  float64
	Email   string
	OrderID string
}

type PaymentOutput struct {
	ResultOrderID   string
	ResultPaymentID string
	CheckoutURL     string
	QRcode          string
	QRCodeB64       string
}
