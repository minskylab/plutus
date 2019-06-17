package plutus

import (
	"errors"
	"time"
)

// SaleState represents the state of any sale
type SaleState string

// Null is a null sale, only for debug porpuses
var Null SaleState = "null"

// Draft is the earliest state of sale
var Draft SaleState = "draft"

// Signed is a signed and final modeling sale
var Signed SaleState = "signed"

// PaidOut is a charged sale
var PaidOut SaleState = "pay_out"

// Done is a done sale
var Done SaleState = "done"

// Sale is a bidirectional transaction
type Sale struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	State         SaleState
	Customer      *Customer
	Products      []Product
	DiscountCodes []DiscountCode
	CardToken     *CardToken
	Charge        *ChargeToken
	CurrencyToPay *Currency
}

// FillID fills the id
func (sale *Sale) FillID() {
	sale.ID = ids.New()
}

func newBasicSale(customerEmail string, products []Product) (*Sale, error) {
	if len(products) == 0 {
		return nil, errors.New("invalid products size, you need at least one product to make a sale")
	}

	currency := products[0].Cost.Currency

	for _, p := range products {
		if p.Cost.Currency != currency {
			return nil, errors.New("invalid currencies, please homologeus your product currencies")
		}
	}

	return &Sale{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		State:         Draft,
		Products:      products,
		DiscountCodes: []DiscountCode{},
		Customer: &Customer{
			Email: customerEmail,
		},
		CurrencyToPay: currency,
	}, nil
}
