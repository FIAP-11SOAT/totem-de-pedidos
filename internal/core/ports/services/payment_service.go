package services

type PaymentService interface {
	CreateWithQRcode() (string, error)
}
