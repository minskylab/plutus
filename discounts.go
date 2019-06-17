package plutus

import "time"

// DiscountType is a type of discount, actually plutus support static or percentual types
type DiscountType string

// StaticDiscount is a static discount (e.g. 100PEN)
var StaticDiscount DiscountType = "static"

// PercentDiscount is a static discount (e.g. 20%)
var PercentDiscount DiscountType = "percent"

// Discount is a discount value
type Discount struct {
	Type     DiscountType
	Percent  float64
	Amount   float64
	Currency *Currency
}

// DiscountUseRecord is a record of discount code use
type DiscountUseRecord struct {
	At time.Time
	By *Customer
}

// DiscountCode represents a discount promotional code
type DiscountCode struct {
	ID      string
	Start   time.Time
	End     time.Time
	MaxUses int
	Uses    []DiscountUseRecord
	Code    string
	Value   *Discount
}

// NewDiscountCode is a wrapper to create a new discount code
type NewDiscountCode struct {
	Code    string
	Start   *time.Time
	End     time.Time
	MaxUses int
	Value   Discount
}

// FillID fills the id
func (d *DiscountCode) FillID() {
	d.ID = ids.New()
}
