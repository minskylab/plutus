package culqi

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/bregydoc/plutus"

	"net/http"
)

const baseTokensURL = "https://secure.culqi.com/tokens/"

type tokenPayload struct {
	PublicKey       string                 `json:"public_key"`
	Email           string                 `json:"email"`
	CardNumber      string                 `json:"card_number"`
	CVV             string                 `json:"cvv"`
	ExpirationYear  int                    `json:"expiration_year"`
	ExpirationMonth int                    `json:"expiration_month"`
	Fingerprint     string                 `json:"fingerprint"`
	Metadata        map[string]interface{} `json:"metadata"`
}

// Token ...
type Token struct {
	Value string `json:"id"`
}

type culqiError struct {
	Object  string `json:"object"`
	Message string `json:"merchant_message"`
}

func (c *Bridge) getNewToken(payload tokenPayload, sessionID string) (*Token, error) {
	buff := bytes.NewBufferString("")
	err := json.NewEncoder(buff).Encode(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", baseTokensURL, buff)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.publicKey)
	req.Header.Set("X-API-VERSION", c.apiVersion)
	req.Header.Set("X-CULQI-ENV", c.env)
	req.Header.Set("X-CULQI-SESS", sessionID)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		returnedError := new(culqiError)
		err = json.NewDecoder(res.Body).Decode(returnedError)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(returnedError.Message)
	}

	token := new(Token)
	err = json.NewDecoder(res.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil

}

func (c *Bridge) generateNewOneUseToken(cardDetails plutus.CardDetails) (*Token, error) {
	sess, err := c.getNewSessionID()
	if err != nil {
		return nil, err
	}

	if cardDetails.Customer == nil {
		return nil, errors.New("please include the customer's email in your card details")
	}
	return c.getNewToken(tokenPayload{
		Email:           cardDetails.Customer.Email,
		CardNumber:      cardDetails.Number,
		CVV:             cardDetails.CVV,
		ExpirationYear:  cardDetails.Expiration.Year,
		ExpirationMonth: cardDetails.Expiration.Month,
		Fingerprint:     "",
		Metadata:        map[string]interface{}{"installments": ""},
	}, sess)
}
