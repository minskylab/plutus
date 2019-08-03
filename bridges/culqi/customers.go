package culqi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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

type customerCreateResponse struct {
	ID string `json:"id"`
}

func (bridge *Bridge) saveCustomerTokenWithEmail(email, customerToken string) error {
	return bridge.repo.Set("customers", email, customerToken)
}

func (bridge *Bridge) getCustomerTokenIfExist(email string) string {
	var customerToken string
	err := bridge.repo.Get("customers", email, &customerToken)
	if err != nil {
		return ""
	}
	return customerToken
}

func (bridge *Bridge) createCustomer(info *CustomerInfo) (string, error) {
	url := baseURL + "/customers"
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(info)
	if err != nil {
		return "", nil
	}

	customerToken := bridge.getCustomerTokenIfExist(info.Email)
	if customerToken != "" {
		return customerToken, nil
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bridge.secretKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 {
		returnedError := new(culqiError)
		err = json.NewDecoder(resp.Body).Decode(returnedError)
		if err != nil {
			return "", err
		}
		return "", errors.New(returnedError.Message)
	}

	customer := new(customerCreateResponse)
	err = json.NewDecoder(resp.Body).Decode(customer)
	if err != nil {
		return "", err
	}

	err = bridge.saveCustomerTokenWithEmail(info.Email, customer.ID)
	if err != nil {
		return "", err
	}

	return customer.ID, nil
}
