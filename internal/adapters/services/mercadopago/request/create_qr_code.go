package request

type Items struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure"`
	UnitPrice   string `json:"unit_price"`
	TotalAmount string `json:"total_amount"`
}

type DynamicQRRequest struct {
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	ExternalReference string  `json:"external_reference"`
	NotificationURL   string  `json:"notification_url"`
	TotalAmount       float64 `json:"total_amount"`
	Items             []Items `json:"items"`
}

type CreateDynamicQRInput struct {
	UserID        string
	ExternalPosID string
	Payload       DynamicQRRequest
}
