package plutus

import "time"

// ChargeParams represents a minimal params to make a charge with your bridge
type ChargeParams struct {
	Amount    float64
	Email     string
	Details   string
	Currency  *Currency
	ExtraInfo *Customer
}

// ChargeToken is the reponse from your payment bridge
type ChargeToken struct {
	ID            string
	Value         string
	Message       string
	WithCardToken CardToken
	CreatedAt     time.Time
}

// FillID fills the id
func (token *ChargeToken) FillID() {
	token.ID = ids.New()
}
