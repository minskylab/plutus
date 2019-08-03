package culqi

import (
	"errors"
	"strings"

	"github.com/bregydoc/plutus"
)

// tokenID: token generaed by culqi, from frontend or plutus ;)
func (q *Bridge) generateNewRecurrentToken(token string, details plutus.CardDetails) (*Card, error) {
	if details.Customer == nil {
		return nil, errors.New("if you want to create a recurrent token, you need to include all customer details")
	}

	if details.Customer.Location == nil {
		return nil, errors.New("please include location details of your customer")
	}

	chunks := strings.Split(details.Customer.Name, " ")

	var firstName, lastName = chunks[0], chunks[1]

	cCode := "PE"
	if details.Customer.Location.CountryCode != "" {
		cCode = details.Customer.Location.CountryCode
	}

	if details.Customer.Location.ZIP != "" && len(details.Customer.Location.ZIP) == 2 {
		cCode = details.Customer.Location.ZIP
	}

	customerID, err := q.createCustomer(&CustomerInfo{
		FirstName:   firstName,
		LastName:    lastName,
		Address:     details.Customer.Location.Address,
		AddressCity: details.Customer.Location.City,
		CountryCode: cCode,
		Email:       details.Customer.Email,
		PhoneNumber: details.Customer.Phone,
	})

	if err != nil {
		return nil, err
	}

	card, err := q.createCard(customerID, token)
	if err != nil {
		return nil, err
	}

	return card, nil
}
