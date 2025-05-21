package mercadopago

type DataID struct {
	ID string `json:"id"`
}

type WebhookPayload struct {
	ID          string `json:"id,omitempty"`
	Data        DataID `json:"data,omitempty"`
	Type        string `json:"type,omitempty"`
	Action      string `json:"action,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	ApiVersion  string `json:"api_version,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	LiveMode    bool   `json:"live_mode,omitempty"`
}
