package repositories

type Payments interface {
	GetPaymentByID(paymentID string) (string, error)
}
