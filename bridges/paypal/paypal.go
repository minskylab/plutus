package paypal

import (
	"errors"
	"net/http"
	"strconv"
	"time"

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

// NewBridge creates a new paypal bridge
func NewBridge(publicKey, privateKey string) (*Bridge, error) {
	bridge := &Bridge{
		publicKey:      publicKey,
		privateKey:     privateKey,
		paypalOAuth:    "https://api.sandbox.paypal.com/v1/oauth2/token/",
		paypalOrderAPI: "https://api.sandbox.paypal.com/v2/checkout/orders/",
	}

	token, err := bridge.genAccessToken()
	if err != nil {
		return nil, err
	}

	bridge.AccessToken = token.AccessToken

	return bridge, nil
}

// Description implements a plutus bridge interface
func (bridge *Bridge) Description() *plutus.BridgeDescription {
	return paypalDescription
}

// NewToken implements a plutus bridge interface
func (bridge *Bridge) NewToken(details plutus.CardDetails, kind plutus.CardTokenType) (*plutus.CardToken, error) {
	return nil, errors.New("invalid operation, currently Paypal Bridge cannot generate card tokens")
}

// MakeCharge implements a plutus bridge interface
func (bridge *Bridge) MakeCharge(source plutus.CardToken, params plutus.ChargeParams) (*plutus.ChargeToken, error) {
	// virtual order ID from generate card token
	orderID := source.Value

	req, err := http.NewRequest(http.MethodGet, bridge.paypalOrderAPI+orderID, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+bridge.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// log.Println("RES CODE:", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		paypalErr := new(paypalError)
		err = json.NewDecoder(res.Body).Decode(paypalErr)
		if err != nil {
			return nil, err
		}

		// * Refresh token
		if paypalErr.Error == "invalid_token" {
			token, err := bridge.genAccessToken()
			if err != nil {
				return nil, err
			}

			bridge.AccessToken = token.AccessToken

			return bridge.MakeCharge(source, params)
		}

		return nil, fmt.Errorf("[from paypal] %s", paypalErr.Error)
	}

	response := new(orderResponse)

	err = json.NewDecoder(res.Body).Decode(&response)

	// log.Printf("%+v\n", response)

	if params.Currency == nil {
		return nil, fmt.Errorf("currency is necessary")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[from paypal] %+v", response)
	}

	if response.Intent != "CAPTURE" {
		return nil, fmt.Errorf("[from paypal] %s", "invalid intent, it should be 'CAPTURE'")
	}

	if len(response.PurchaseUnits) == 0 {
		return nil, fmt.Errorf("[from paypal] %s", "invalid purchase units, it should have a more length")
	}

	amount := response.PurchaseUnits[0].Amount

	amountFloat, err := strconv.ParseFloat(amount.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("[from paypal] %s", "invalid amount number: "+amount.Value)
	}

	if params.Amount != amountFloat {
		return nil, fmt.Errorf("invalid amount, your order did capture %.2f and you pass %.2f", amountFloat, params.Amount)
	}

	if params.Currency.Code != amount.CurrencyCode {
		return nil, fmt.Errorf("[from paypal] no merge between %s and %s currencies", params.Currency.Code, amount.CurrencyCode)
	}

	if params.Email != response.Payer.EmailAddress {
		return nil, fmt.Errorf("[from paypal] your payer is %s or %s?", params.Email, response.Payer.EmailAddress)
	}

	if response.Status != "COMPLETED" {
		return nil, fmt.Errorf("[from paypal] your order are not complete")
	}
	// TODO: Verify the time differencies

	charge := &plutus.ChargeToken{
		CreatedAt:     time.Now(),
		Value:         response.ID,
		Message:       response.Intent + " " + response.Status,
		WithCardToken: source,
	}

	return charge, nil

}

// MakeRefund implements a plutus bridge interface
func (bridge *Bridge) MakeRefund(source plutus.ChargeToken, params plutus.RefundParams) (*plutus.RefundToken, error) {
	return nil, errors.New("invalid operation, currently Paypal Bridge can not makes refunds")
}
