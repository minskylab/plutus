package plutus

import (
	"context"
	"errors"
	"time"

	plutus "github.com/bregydoc/plutus/proto"
)

// NewCardToken implements a grpc plutus service
func (e *SalesEngine) NewCardToken(c context.Context, p *plutus.NewCardTokenRequest) (*plutus.CardToken, error) {
	var customer *Customer
	if p.Customer != nil {
		var location *Location
		if p.Customer.Location != nil {
			location = &Location{
				Address: p.Customer.Location.Address,
				City:    p.Customer.Location.City,
				ZIP:     p.Customer.Location.Zip,
				State:   p.Customer.Location.State,
			}
		}
		customer = &Customer{
			ID:       p.Customer.Id,
			Email:    p.Customer.Email,
			Name:     p.Customer.Name,
			Person:   p.Customer.Person,
			Phone:    p.Customer.Phone,
			Location: location,
		}
	}

	token, err := e.Bridge.NewToken(CardDetails{
		Number: p.Card.Number,
		Expiration: Date{
			Month: int(p.Card.ExpMont),
			Year:  int(p.Card.ExpYear),
		},
		CVV:      p.Card.Cvc,
		Customer: customer,
	}, cardTokenTypeFromProto(p.Type))
	if err != nil {
		return nil, err
	}

	token, err = e.Repository.SaveCardToken(token)
	if err != nil {
		return nil, err
	}

	return cardTokenToProto(*token), nil
}

// NewCardTokenFromNative implements a grpc plutus service
func (e *SalesEngine) NewCardTokenFromNative(c context.Context, p *plutus.NewCardTokenNativeRequest) (*plutus.CardToken, error) {
	nativeToken := p.Token
	t := cardTokenTypeFromProto(p.Type)
	var customer *Customer
	if p.Customer != nil {
		c := customerFromProto(p.Customer)
		customer = &c
	}

	token := &CardToken{
		CreatedAt: time.Now(),
		Type:      t,
		Value:     nativeToken,
		WithCard: EncodedCardDetails{
			Customer: customer,
		},
	}

	var err error
	token, err = e.Repository.SaveCardToken(token)
	if err != nil {
		return nil, err
	}

	return cardTokenToProto(*token), nil
}

// GetCardTokenOfCustomerByID implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByID(c context.Context, p *plutus.CardTokenByID) (*plutus.CardToken, error) {
	token, err := e.Repository.GetCardToken(p.Id)
	if err != nil {
		return nil, err
	}

	return cardTokenToProto(*token), nil
}

// GetCardTokenOfCustomerByCustomer implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByCustomer(c context.Context, p *plutus.CardTokenByCustomer) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// DeleteCardToken implements a grpc plutus service
func (e *SalesEngine) DeleteCardToken(c context.Context, p *plutus.CardTokenByID) (*plutus.CardToken, error) {
	token, err := e.Repository.RemoveCardToken(p.Id)
	if err != nil {
		return nil, err
	}

	return cardTokenToProto(*token), nil
}

// NewFastSale implements a grpc plutus service
func (e *SalesEngine) NewFastSale(c context.Context, p *plutus.FastSale) (*plutus.Sale, error) {
	sale, err := newBasicSale(p.CustomerEmail, productsFromProto(p.Products))
	if err != nil {
		return nil, err
	}

	sale, err = e.Repository.SaveSale(sale)
	if err != nil {
		return nil, err
	}

	return saleToProto(*sale), nil
}

// NewSale implements a grpc plutus service
func (e *SalesEngine) NewSale(c context.Context, p *plutus.NewSaleRequest) (*plutus.Sale, error) {
	customer := customerFromProto(p.Customer)
	sale := &Sale{
		Customer:  &customer,
		Products:  productsFromProto(p.Products),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		State:     Draft,
	}

	var err error
	sale, err = e.Repository.SaveSale(sale)
	if err != nil {
		return nil, err
	}

	return saleToProto(*sale), nil

}

// GetSale implements a grpc plutus service
func (e *SalesEngine) GetSale(c context.Context, p *plutus.SaleIDRequest) (*plutus.Sale, error) {
	sale, err := e.Repository.GetSale(p.Id)
	if err != nil {
		return nil, err
	}

	return saleToProto(*sale), nil
}

// UpdateSale implements a grpc plutus service
func (e *SalesEngine) UpdateSale(c context.Context, p *plutus.SaleUpdateRequest) (*plutus.Sale, error) {
	sale, err := e.Repository.UpdateSale(p.Id, saleFromProto(p.UpdateData))
	if err != nil {
		return nil, err
	}

	return saleToProto(*sale), nil
}

// DeliverSaleReceipt implements a grpc plutus service
func (e *SalesEngine) DeliverSaleReceipt(c context.Context, p *plutus.DeliverSaleRequest) (*plutus.DeliverChannelResponse, error) {
	panic("unimplemented")
}

// ChargeSaleByID implements a grpc plutus service
func (e *SalesEngine) ChargeSaleByID(c context.Context, p *plutus.ChargeSaleRequest) (*plutus.ChargeToken, error) {
	sale, err := e.Repository.GetSale(p.SaleID)
	if err != nil {
		return nil, err
	}

	cardToken, err := e.Repository.GetCardToken(p.CardTokenID)
	if err != nil {
		return nil, err
	}

	customer, _ := e.Repository.GetCustomer(sale.Customer.ID)
	if customer != nil {
		sale.Customer = customer
	}

	currency := sale.CurrencyToPay
	total := 0.0
	for _, product := range sale.Products {
		if currency.Name != product.Cost.Currency.Name {
			return nil, errors.New("incompatible currencies, actually plutus only accepts homologous currency")
		}
		total += product.Cost.Amount
	}

	details := "plutus sale charge"
	if p.Details != "" {
		details = p.Details
	}

	charge, err := e.Bridge.MakeCharge(*cardToken, ChargeParams{
		Amount:    total,
		Currency:  currency,
		Details:   details,
		Email:     sale.Customer.Email,
		ExtraInfo: sale.Customer,
	})

	if err != nil {
		return nil, err
	}

	sale.CardToken = cardToken
	sale.Charge = charge
	sale.UpdatedAt = time.Now()
	sale.State = PaidOut

	sale, err = e.Repository.UpdateSale(sale.ID, *sale)
	if err != nil {
		return nil, err
	}

	return chargeTokenToProto(*charge), nil

}

// ChargeSaleWithNativeToken implements a grpc plutus service
func (e *SalesEngine) ChargeSaleWithNativeToken(c context.Context, p *plutus.ChargeWithNativeToken) (*plutus.ChargeToken, error) {
	sale, err := e.Repository.GetSale(p.SaleID)
	if err != nil {
		return nil, err
	}

	customer, _ := e.Repository.GetCustomer(sale.Customer.ID)
	if customer != nil {
		sale.Customer = customer
	}

	currency := sale.CurrencyToPay
	total := 0.0
	for _, product := range sale.Products {
		if currency.Name != product.Cost.Currency.Name {
			return nil, errors.New("incompatible currencies, actually plutus only accepts homologous currency")
		}
		total += product.Cost.Amount
	}

	details := "plutus sale charge"
	if p.Details != "" {
		details = p.Details
	}

	cardToken := &CardToken{
		CreatedAt: time.Now(),
		Type:      OneUseToken,
		Value:     p.NativeToken,
		WithCard: EncodedCardDetails{
			Customer: customer,
		},
	}

	cardToken, err = e.Repository.SaveCardToken(cardToken)
	if err != nil {
		return nil, err
	}

	charge, err := e.Bridge.MakeCharge(*cardToken, ChargeParams{
		Amount:    total,
		Currency:  currency,
		Details:   details,
		Email:     sale.Customer.Email,
		ExtraInfo: sale.Customer,
	})

	if err != nil {
		return nil, err
	}

	sale.CardToken = cardToken
	sale.Charge = charge
	sale.UpdatedAt = time.Now()
	sale.State = PaidOut

	sale, err = e.Repository.UpdateSale(sale.ID, *sale)
	if err != nil {
		return nil, err
	}

	return chargeTokenToProto(*charge), nil

}

// DoneSale implements a grpc plutus service
func (e *SalesEngine) DoneSale(c context.Context, p *plutus.SaleIDRequest) (*plutus.Sale, error) {
	sale, err := e.Repository.GetSale(p.Id)
	if err != nil {
		return nil, err
	}

	sale.State = Done

	sale, err = e.Repository.UpdateSale(sale.ID, *sale)
	if err != nil {
		return nil, err
	}

	return saleToProto(*sale), nil
}

// CreateDiscountCode implements a grpc plutus service
func (e *SalesEngine) CreateDiscountCode(c context.Context, p *plutus.DiscountCodeRequest) (*plutus.DiscountCode, error) {
	panic("unimplemented")
}

// GetDiscountCode implements a grpc plutus service
func (e *SalesEngine) GetDiscountCode(c context.Context, p *plutus.DiscountCodeID) (*plutus.DiscountCode, error) {
	panic("unimplemented")
}

// ValidateDiscountCode implements a grpc plutus service
func (e *SalesEngine) ValidateDiscountCode(c context.Context, p *plutus.DiscountCodeValue) (*plutus.DiscountCodeExist, error) {
	panic("unimplemented")
}

// GetActiveDiscountCodes implements a grpc plutus service
func (e *SalesEngine) GetActiveDiscountCodes(c context.Context, p *plutus.ActiveDiscountsRequest) (*plutus.DiscountCodes, error) {
	panic("unimplemented")
}

// DeleteDiscountCode implements a grpc plutus service
func (e *SalesEngine) DeleteDiscountCode(c context.Context, p *plutus.DiscountCodeID) (*plutus.DiscountCodes, error) {
	panic("unimplemented")
}
