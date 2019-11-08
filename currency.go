package plutus

// Currency is struct to describe a currency
type Currency struct {
	Name   string
	Code   string
	Symbol string
}

// PEN represents a Peruvian Currency
var PEN = &Currency{
	Name:   "PEN",
	Code:   "PEN",
	Symbol: "s/",
}

// USD represents an American Dollar
var USD = &Currency{
	Name:   "USD",
	Code:   "USD",
	Symbol: "$",
}

// AvailableCurrencies is a map with available currencies
var AvailableCurrencies = map[string]*Currency{
	PEN.Code: PEN,
	USD.Code: USD,
}
