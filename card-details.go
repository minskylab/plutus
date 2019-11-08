package plutus

import (
	"fmt"
	"github.com/bregydoc/plutus/creditcard"
)

// Date is a simple wrap for year and month date
type Date struct {
	Year  int
	Month int
}

// CardDetails represents a minimal details to encode your card sensible deatils
type CardDetails struct {
	Number     string
	Expiration Date
	CVV        string
	Customer   *Customer
}

// EncodedCardDetails represents a encoded card details (hidden complete number and expiration year too)
type EncodedCardDetails struct {
	Number         string
	ExpirationYear int
	Customer       *Customer
}

func (card *CardDetails) Validate(mods ...string) bool {
	c := creditcard.Card{
		Number: card.Number,
		Cvv:    card.CVV,
		Month:  fmt.Sprintf("%d", card.Expiration.Month),
		Year:   fmt.Sprintf("%d", card.Expiration.Year),
	}
	if err := c.Method(); err != nil {
		return false
	}

	if err := c.Validate(true); err != nil {
		return false
	}

	return true
}
