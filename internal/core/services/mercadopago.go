package services

// interface que chama o mercado pago
type MercadoPago struct {
}

func New() MercadoPago {
	return MercadoPago{}
}

func (m *MercadoPago) Call() error {
	return nil
}
