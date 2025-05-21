package entity

import "time"

type Payment struct {
	ID          int
	Amount      float64
	PaymentDate time.Time
	Status      string
	Provider    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
