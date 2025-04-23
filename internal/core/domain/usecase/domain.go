package domain

type OrderUsecase struct {
}

func New() *OrderUsecase {
	return &OrderUsecase{}
}

func (o *OrderUsecase) CreateOrder() error {
	return nil
}
