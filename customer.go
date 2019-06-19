package plutus

// Location is a set of fields to localize to customer
type Location struct {
	Address     string
	City        string
	State       string
	CountryCode string
	ZIP         string
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

// FillID fills the id
func (customer *Customer) FillID() {
	customer.ID = ids.New()
}
