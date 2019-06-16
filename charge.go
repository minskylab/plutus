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
	ID        string
	Value     string
	Message   string
	CreatedAt time.Time
}

func (token *ChargeToken) fillID() {
	token.ID = ids.New()
}
