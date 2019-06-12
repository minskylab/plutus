package plutus

// Location is a set of fields to localize to customer
type Location struct {
	Address string
	City    string
	State   string
	ZIP     string
}

// Customer represents a customer minimal information
type Customer struct {
	ID       string
	Person   string
	Name     string
	Email    string
	Phone    string
	Location *Location
}
