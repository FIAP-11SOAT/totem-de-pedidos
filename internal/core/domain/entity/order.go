package entity

import "time"

type Order struct {
	ID                   int
	OrderDate            time.Time
	NotificationAttempts int
	Status               string
	TotalAmount          float64
	CreatedAt            time.Time
	CustomerID           int
	Items                []OrderItem
}

type OrderItem struct {
	ID        int
	Quantity  int
	Price     float64
	CreatedAt time.Time
	OrderID   int
	ProductID int
}
