package plutus

// Date is a simple wrap for year and month date
type Date struct {
	Year  int
	Month int
}

// Location is a set of fields to localize to customer
type Location struct {
	Address string
	City    string
	State   string
	ZIP     string
}

// Customer represents a customer minimal information
type Customer struct {
	Name     string
	Email    string
	Phone    string
	Location *Location
}

// CardDetails represents a minimal details to encode your card sensible deatils
type CardDetails struct {
	Number     string
	Expiration Date
	CVV        string
}
