package output

import "time"

type PaymentOutput struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
	Provider    string    `json:"provider"`
}
