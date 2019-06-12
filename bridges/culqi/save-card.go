package culqi

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bregydoc/plutus"
)

// tokenID: token generaed by culqi, from frontend or plutus ;)
func (q *PlutusBridge) generateNewRecurrentToken(token string, details plutus.CardDetails) (*Card, error) {
	if details.Customer == nil {
		return nil, errors.New("if you want to create a recurrent token, you need to include all customer details")
	}

	if details.Customer.Location == nil {
		return nil, errors.New("please include location details of your customer")
	}

	chunks := strings.Split(details.Customer.Name, " ")

	var firstName, lastName = chunks[0], chunks[1]

	customerID, err := q.createCustomer(&CustomerInfo{
		FirstName:   firstName,
		LastName:    lastName,
		Address:     details.Customer.Location.Address,
		AddressCity: details.Customer.Location.City,
		CountryCode: details.Customer.Location.ZIP,
		Email:       details.Customer.Email,
		PhoneNumber: details.Customer.Phone,
	})

	if err != nil {
		return nil, fmt.Errorf("[from culqi] %s", err.Error())
	}

	card, err := q.createCard(customerID, token)
	if err != nil {
		return nil, fmt.Errorf("[from culqi] %s", err.Error())
	}

	return card, nil
}
