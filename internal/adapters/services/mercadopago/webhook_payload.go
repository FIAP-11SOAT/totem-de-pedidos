package mercadopago

type DataID struct {
	ID string `json:"id"`
}

type WebhookPayload struct {
	ID          string `json:"id"`
	Data        DataID `json:"data"`
	Type        string `json:"type"`
	Action      string `json:"action"`
	UserID      int64  `json:"user_id"`
	ApiVersion  string `json:"api_version"`
	DateCreated string `json:"date_created"`
	LiveMode    bool   `json:"live_mode"`
}
