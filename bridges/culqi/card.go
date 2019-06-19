package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.culqi.com/v2"

// Card represents the culqi response to /card endpoint
type Card struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Token      struct {
		ID         string `json:"id"`
		CardNumber string `json:"card_number"`
		Active     bool
		Iin        struct {
			CardBrand    string `json:"card_brand"`
			CardType     string `json:"card_type"`
			CardCategory string `json:"card_category"`
		} `json:"iin"`
	} `json:"token"`
}

type cardCreateInput struct {
	CustomerID string `json:"customer_id"`
	TokenID    string `json:"token_id"`
}

func (bridge *PlutusBridge) createCard(customerID, tokenID string) (*Card, error) {
	url := baseURL + "/cards"
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(cardCreateInput{
		CustomerID: customerID,
		TokenID:    tokenID,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bridge.secretKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	card := new(Card)
	err = json.NewDecoder(resp.Body).Decode(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (bridge *PlutusBridge) getCard(id string) (*Card, error) {
	url := baseURL + fmt.Sprintf("/cards/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bridge.secretKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	card := new(Card)
	err = json.NewDecoder(resp.Body).Decode(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}
