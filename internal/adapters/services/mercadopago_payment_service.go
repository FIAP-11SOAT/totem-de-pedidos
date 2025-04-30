package services

type MercadoPagoPaymentService struct{}

func NewMercadoPagoPaymentService() *MercadoPagoPaymentService {
	return &MercadoPagoPaymentService{}
}

func (s *MercadoPagoPaymentService) CreateWithQRcode() {}
