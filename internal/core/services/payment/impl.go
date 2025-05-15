package payment

type PaymentService interface {
	CreatePix(CreatePixInput) (CreatePixOutput, error)
}
