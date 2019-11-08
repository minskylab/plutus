package dummy

import (
	"errors"
	"github.com/bregydoc/plutus"
	"log"
	"strconv"
	"time"
	"math/rand"
)
var r = rand.New(rand.NewSource(time.Now().Unix()))


type Bridge struct {}


func (bridge *Bridge) Describe() *plutus.BridgeDescription {
	return dummyDesc
}

func (bridge *Bridge) NewToken(details plutus.CardDetails, kind plutus.CardTokenType) (*plutus.CardToken, error) {
	log.Printf("new token from card: %s, customer: %s", details.Number, details.Customer)
	log.Printf("                     %s", kind)

	if !details.Validate() {
		return nil, errors.New("invalid credit card")
	}
	return &plutus.CardToken{
		ID:        strconv.FormatInt(r.Int63(), 10),
		Value:     "0x001",
		Type:      kind,
		WithCard:  plutus.EncodedCardDetails{
			Number:         "41111*****1111",
			ExpirationYear: 2021,
			Customer:       &plutus.Customer{
				ID:       "",
				Person:   "0x2123",
				Name:     "Dummy Customer",
				Email:    "dummy@dev.to",
				Phone:    "",
				Location: nil,
			},
		},
		CreatedAt: time.Now(),
	}, nil
}

func (bridge *Bridge) MakeCharge(source plutus.CardToken, params plutus.ChargeParams) (*plutus.ChargeToken, error) {
	if source.Value != "0x001" {
		return nil, errors.New("invalid dummy card token")
	}

	if source.WithCard.Customer == nil {
		return nil, errors.New("customer information must be different to nil")
	}

	if params.Email != source.WithCard.Customer.Email {
		return nil, errors.New("the email of your charge must be equals to customer card")
	}

	return &plutus.ChargeToken{
		ID:            strconv.FormatInt(r.Int63(), 10),
		Value: 		   "0x0c1",
		Message:       "I'm dummy",
		WithCardToken: source,
		CreatedAt:     time.Now(),
	}, nil
}

func (bridge *Bridge) MakeRefund(source plutus.ChargeToken, params plutus.RefundParams) (*plutus.RefundToken, error) {
	panic("unimplemented")
}
