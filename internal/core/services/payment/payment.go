package payment

type paymentService struct{}

func New() *paymentService { return &paymentService{} }

func (p *paymentService) Payment(_ PaymentInput) (PaymentOutput, error) {
	return PaymentOutput{}, nil
}
