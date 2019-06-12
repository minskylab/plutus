package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.culqi.com/v2"

// CustomerInfo is an customer object of Culqi Service
type CustomerInfo struct {
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	Address     string            `json:"address"`
	AddressCity string            `json:"address_city"`
	CountryCode string            `json:"country_code"`
	PhoneNumber string            `json:"phone_number"`
	Metadata    map[string]string `json:"metadata"`
}

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

type customerCreateResponse struct {
	ID string `json:"id"`
}

type cardCreateInput struct {
	CustomerID string `json:"customer_id"`
	TokenID    string `json:"token_id"`
}

func (q *PlutusBridge) createCustomer(info *CustomerInfo) (string, error) {
	url := baseURL + "/customers"
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(info)
	if err != nil {
		return "", nil
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", q.secretKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	customer := new(customerCreateResponse)
	err = json.NewDecoder(resp.Body).Decode(customer)
	if err != nil {
		return "", err
	}

	return customer.ID, nil
}

func (q *PlutusBridge) createCard(customerID, tokenID string) (*Card, error) {
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", q.secretKey))

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

func (q *PlutusBridge) getCard(id string) (*Card, error) {
	url := baseURL + fmt.Sprintf("/cards/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", q.secretKey))

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
