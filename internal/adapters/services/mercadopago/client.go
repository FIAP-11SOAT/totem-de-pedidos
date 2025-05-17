package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago/request"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/services/mercadopago/response"
)

var client = &http.Client{}

type Store struct {
	AccessToken string
}

func (s *Store) CreateDynamicQR(input *request.CreateDynamicQRInput) (response.DynamicQRResponse, error) {
	endpoint := fmt.Sprintf(
		"https://api.mercadopago.com/instore/orders/qr/seller/collectors/%s/pos/%s/qrs",
		input.UserID,
		input.ExternalPosID,
	)

	respObj := response.DynamicQRResponse{}
	content, err := json.Marshal(input.Payload)
	if err != nil {
		return respObj, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(content))
	if err != nil {
		return respObj, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return respObj, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return respObj, err
	}

	return respObj, nil
}
