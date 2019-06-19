package culqi

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/asdine/storm"

	"github.com/bregydoc/plutus"
)

// PlutusBridge represents a bridge between your culqi service and your plutus service
type PlutusBridge struct {
	publicKey  string
	secretKey  string
	env        string
	apiVersion string
	repo       *storm.DB
}

// NewPlutusBridge returns a new culqi bridge instance
func NewPlutusBridge(publicKey, secretKey string) (*PlutusBridge, error) {
	var err error
	repo, err := storm.Open("culqi-helper.db")
	if err != nil {
		return nil, err
	}
	env := "prod"
	if strings.Contains(secretKey, "test") {
		env = "test"
	}
	return &PlutusBridge{
		publicKey:  publicKey,
		secretKey:  secretKey,
		env:        env,
		apiVersion: "2",
		repo:       repo,
	}, nil
}

// NewToken returns a new token of your card, that is an implementation of plutus bridge
func (bridge *PlutusBridge) NewToken(details plutus.CardDetails, kind plutus.CardTokenType) (*plutus.CardToken, error) {
	if err := validateCardDetails(details); err != nil {
		return nil, err
	}

	switch kind {
	case plutus.OneUseToken:
		token, err := bridge.generateNewOneUseToken(details)
		if err != nil {
			return nil, fmt.Errorf("[from culqi] %s", err.Error())
		}

		encodedNumberCard := ""
		if len(details.Number) >= 4 {
			encodedNumberCard = details.Number[len(details.Number)-4:]
			encodedNumberCard = strings.Repeat("*", len(details.Number)-4) + encodedNumberCard
		}

		return &plutus.CardToken{
			CreatedAt: time.Now(),
			Type:      kind,
			Value:     token.Value,
			WithCard: plutus.EncodedCardDetails{
				Number:         encodedNumberCard,
				Customer:       details.Customer,
				ExpirationYear: details.Expiration.Year,
			},
		}, nil

	case plutus.RecurrentToken:
		token, err := bridge.generateNewOneUseToken(details)
		if err != nil {
			return nil, fmt.Errorf("[from culqi] %s", err.Error())
		}

		card, err := bridge.generateNewRecurrentToken(token.Value, details)
		if err != nil {
			return nil, fmt.Errorf("[from culqi] %s", err.Error())
		}

		encodedNumberCard := ""
		if len(details.Number) >= 4 {
			encodedNumberCard = details.Number[len(details.Number)-4:]
			encodedNumberCard = strings.Repeat("*", len(details.Number)-4) + encodedNumberCard
		}

		return &plutus.CardToken{
			CreatedAt: time.Now(),
			Type:      kind,
			Value:     card.ID,
			WithCard: plutus.EncodedCardDetails{
				Number:         encodedNumberCard,
				Customer:       details.Customer,
				ExpirationYear: details.Expiration.Year,
			},
		}, nil
	default:
		break
	}

	return nil, errors.New("invalid token type")

}

// MakeCharge make a charge with culqi, that implements plutus bridge
func (bridge *PlutusBridge) MakeCharge(source plutus.CardToken, params plutus.ChargeParams) (*plutus.ChargeToken, error) {
	return bridge.executeCharge(source, params)
}

// MakeRefund ...
func (bridge *PlutusBridge) MakeRefund(source plutus.ChargeToken, params plutus.RefundParams) (*plutus.RefundToken, error) {
	return nil, errors.New("refund not implemented")
}
