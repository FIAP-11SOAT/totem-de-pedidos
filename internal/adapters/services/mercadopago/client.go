package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
}

type DynamicQRResponse struct{}

type CreateDynamicQRInput struct {
	UserID        string
	ExternalPosID string
	Payload       DynamicQRRequest
}

func CreateDynamicQR(input CreateDynamicQRInput) (string, error) {
	endpoint := fmt.Sprintf(
		"https://api.mercadopago.com/instore/orders/qr/seller/collectors/%s/pos/%s/qrs",
		input.UserID,
		input.ExternalPosID,
	)

	content, err := json.Marshal(input.Payload)
	if err != nil {
		return "", err
	}

	// #nosec G107
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(content))
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create dynamic QR code: %s", resp.Status)
	}

	return resp.Request.URL.String(), nil
}
