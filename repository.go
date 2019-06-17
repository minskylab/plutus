package plutus

// Repository represents a bag where you can to put your basic objects
type Repository interface {
	SaveCustomer(customer *Customer) (*Customer, error)
	GetCustomer(ID string) (*Customer, error)
	UpdateCustomer(ID string, updatePayload Customer) (*Customer, error)
	RemoveCustomer(ID string) (*Customer, error)

	SaveCardToken(cardToken *CardToken) (*CardToken, error)
	GetCardToken(ID string) (*CardToken, error)
	UpdateCardToken(ID string, updatePayload CardToken) (*CardToken, error)
	RemoveCardToken(ID string) (*CardToken, error)

	SaveChargeToken(chargeToken *ChargeToken) (*ChargeToken, error)
	GetChargeToken(ID string) (*ChargeToken, error)
	UpdateChargeToken(ID string, updatePayload ChargeToken) (*ChargeToken, error)
	RemoveChargeToken(ID string) (*ChargeToken, error)

	SaveSale(sale *Sale) (*Sale, error)
	GetSale(ID string) (*Sale, error)
	UpdateSale(ID string, updatePayload Sale) (*Sale, error)
	RemoveSale(ID string) (*Sale, error)
}
