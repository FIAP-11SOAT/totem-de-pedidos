package input

import "errors"

type OrderInput struct {
	CustomerID int              `json:"customer_id"`
	Items      []OrderItemInput `json:"items"`
}

type OrderItemInput struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (o *OrderInput) Validate() error {
	return nil
}

type UpdateOrderInput struct {
	Status string `json:"status"`
}

func (u *UpdateOrderInput) Validate() error {
	if u.Status == "" {
		return errors.New("status is required")
	}
	return nil
}

type OrderFilterInput struct {
	ID                   *int   `query:"id"`
	Status               string `query:"status"`
	CustomerID           *int   `query:"customer_id"`
	NotificationAttempts *int   `query:"notification_attempts"`
}
