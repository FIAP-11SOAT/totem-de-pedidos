package payment

type CreatePixInput struct {
	Amount  float64
	Email   string
	OrderID string
}

type CreatePixOutput struct {
	ResultOrderID   string
	ResultPaymentID string
	CheckoutURL     string
	QRcode          string
	QRCodeB64       string
}
