package plutus

// SupportContact wraps the support contact of your company
type SupportContact struct {
	Email string
	Phone string
}

// Company serves to describe your company
type Company struct {
	Name        string
	OfficialWeb string
	Support     SupportContact
	Custom      map[string]interface{}
}
