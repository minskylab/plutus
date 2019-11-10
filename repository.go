package plutus

import "context"

// Repository represents a bag where you can to put your basic objects
type Repository interface {
	SaveCustomer(c context.Context, customer *Customer) (*Customer, error)
	GetCustomer(c context.Context, ID string) (*Customer, error)
	UpdateCustomer(c context.Context, ID string, updatePayload Customer) (*Customer, error)
	RemoveCustomer(c context.Context, ID string) (*Customer, error)

	SaveCardToken(c context.Context, cardToken *CardToken) (*CardToken, error)
	GetCardToken(c context.Context, ID string) (*CardToken, error)
	UpdateCardToken(c context.Context, ID string, updatePayload CardToken) (*CardToken, error)
	RemoveCardToken(c context.Context, ID string) (*CardToken, error)

	SaveChargeToken(c context.Context, chargeToken *ChargeToken) (*ChargeToken, error)
	GetChargeToken(c context.Context, ID string) (*ChargeToken, error)
	UpdateChargeToken(c context.Context, ID string, updatePayload ChargeToken) (*ChargeToken, error)
	RemoveChargeToken(c context.Context, ID string) (*ChargeToken, error)

	SaveSale(c context.Context, sale *Sale) (*Sale, error)
	GetSale(c context.Context, ID string) (*Sale, error)
	UpdateSale(c context.Context, ID string, updatePayload Sale) (*Sale, error)
	RemoveSale(c context.Context, ID string) (*Sale, error)
}
