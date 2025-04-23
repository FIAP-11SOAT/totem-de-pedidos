package ports

type MercadoPagoInterface interface {
	Call() error
}
