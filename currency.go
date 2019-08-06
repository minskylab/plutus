package plutus

// Currency is struct to describe a currency
type Currency struct {
	Name       string
	Symbol     string
	Multiplier int
}

// UNIT is the basic unit of currency
var UNIT = &Currency{
	Multiplier: 1,
	Name:       "UNIT",
	Symbol:     "-",
}

// PEN represents a Peruvian Currency
var PEN = &Currency{
	Multiplier: 100,
	Name:       "PEN",
	Symbol:     "s/",
	// Base:       &UNIT,
}

// USD represents an American Dollar
var USD = &Currency{
	Multiplier: 100,
	Name:       "USD",
	Symbol:     "$",
	// Base:       &UNIT,
}

// AvailableCurrencies is a map with available currencies
var AvailableCurrencies = map[string]*Currency{
	PEN.Name: PEN,
	USD.Name: USD,
}
