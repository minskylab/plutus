package plutus

import "time"

// RefundParams represents a minimal params to make a refund with your bridge
type RefundParams struct {
	Amount   float64
	Reason   string
	Currency *Currency
}

// RefundToken is the reponse from your payment bridge
type RefundToken struct {
	ID        string
	Value     string
	CreatedAt time.Time
}
