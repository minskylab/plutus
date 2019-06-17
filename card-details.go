package plutus

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
