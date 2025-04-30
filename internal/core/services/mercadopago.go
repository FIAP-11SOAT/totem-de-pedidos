package services

type MercadoPago struct {
}

func New() MercadoPago {
	return MercadoPago{}
}

func (m *MercadoPago) Call() error {
	return nil
}
