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
	ID        string `storm:"id"`
	Value     string `storm:"unique"`
	CreatedAt time.Time
}

// FillID fills the id
func (refund *RefundToken) FillID() {
	refund.ID = ids.New()
}
