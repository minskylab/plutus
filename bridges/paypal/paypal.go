package paypal

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/bregydoc/plutus"
)

var paypalDescription = &plutus.BridgeDescription{
	Name:                 "paypal",
	Version:              "0.0.2",
	CanGenerateCardToken: false,
	CanMakeCharge:        true,
	CanMakeRefund:        false,
}

// Bridge enables a new plutus bridge to comunicate your orders with PAYPAL
type Bridge struct {
	publicKey      string
	privateKey     string
	paypalOAuth    string
	paypalOrderAPI string
	AccessToken    string
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewBridge(publicKey, privateKey string) (*Bridge, error) {
	bridge := &Bridge{
		publicKey:      publicKey,
		privateKey:     privateKey,
		paypalOAuth:    "https://api.sandbox.paypal.com/v1/oauth2/token/",
		paypalOrderAPI: "https://api.sandbox.paypal.com/v2/checkout/orders/",
	}

	credentials := fmt.Sprintf("%s:%s", bridge.publicKey, bridge.privateKey)
	crd := base64.StdEncoding.EncodeToString([]byte(credentials))

	body := bytes.NewBufferString("grant_type=client_credentials")
	req, err := http.NewRequest("POST", bridge.paypalOAuth, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+crd)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println("RES CODE:", res.StatusCode)

	token := new(tokenResponse)
	err = json.NewDecoder(res.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	bridge.AccessToken = token.AccessToken

	return bridge, nil
}

func (bridge *Bridge) Describe() *plutus.BridgeDescription {
	return paypalDescription
}

func (bridge *Bridge) NewToken(details plutus.CardDetails, kind plutus.CardTokenType) (*plutus.CardToken, error) {
	return nil, errors.New("invalid operation, currently Paypal Bridge cannot generate card tokens")
}

func (bridge *Bridge) MakeCharge(source plutus.CardToken, params plutus.ChargeParams) (*plutus.ChargeToken, error) {
	panic("unimplemented")
}

func (bridge *Bridge) MakeRefund(source plutus.ChargeToken, params plutus.RefundParams) (*plutus.RefundToken, error) {
	return nil, errors.New("invalid operation, currently Paypal Bridge can not makes refunds")
}
