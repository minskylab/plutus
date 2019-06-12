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
