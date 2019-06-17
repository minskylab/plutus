package plutus

import (
	"context"

	plutus "github.com/bregydoc/plutus/proto"
)

// NewCardToken implements a grpc plutus service
func (e *SalesEngine) NewCardToken(context.Context, *plutus.NewCardTokenRequest) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// GetCardTokenOfCustomerByID implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByID(context.Context, *plutus.CardTokenByID) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// GetCardTokenOfCustomerByCustomer implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByCustomer(context.Context, *plutus.CardTokenByCustomer) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// DeleteCardToken implements a grpc plutus service
func (e *SalesEngine) DeleteCardToken(context.Context, *plutus.CardTokenByID) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// NewFastSale implements a grpc plutus service
func (e *SalesEngine) NewFastSale(context.Context, *plutus.FastSale) (*plutus.Sale, error) {
	panic("unimplemented")
}

// NewSale implements a grpc plutus service
func (e *SalesEngine) NewSale(context.Context, *plutus.NewSaleRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// GetSale implements a grpc plutus service
func (e *SalesEngine) GetSale(context.Context, *plutus.SaleIDRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// UpdateSale implements a grpc plutus service
func (e *SalesEngine) UpdateSale(context.Context, *plutus.SaleUpdateRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// DeliverSaleReceipt implements a grpc plutus service
func (e *SalesEngine) DeliverSaleReceipt(context.Context, *plutus.DeliverSaleRequest) (*plutus.DeliverChannelResponse, error) {
	panic("unimplemented")
}

// ChargeSaleByID implements a grpc plutus service
func (e *SalesEngine) ChargeSaleByID(context.Context, *plutus.SaleIDRequest) (*plutus.ChargeToken, error) {
	panic("unimplemented")
}

// ChargeSaleWithNativeToken implements a grpc plutus service
func (e *SalesEngine) ChargeSaleWithNativeToken(context.Context, *plutus.ChargeWithNativeToken) (*plutus.ChargeToken, error) {
	panic("unimplemented")
}

// DoneSale implements a grpc plutus service
func (e *SalesEngine) DoneSale(context.Context, *plutus.SaleIDRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// CreateDiscountCode implements a grpc plutus service
func (e *SalesEngine) CreateDiscountCode(context.Context, *plutus.DiscountCodeRequest) (*plutus.DiscountCode, error) {
	panic("unimplemented")
}

// GetDiscountCode implements a grpc plutus service
func (e *SalesEngine) GetDiscountCode(context.Context, *plutus.DiscountCodeID) (*plutus.DiscountCode, error) {
	panic("unimplemented")
}

// ValidateDiscountCode implements a grpc plutus service
func (e *SalesEngine) ValidateDiscountCode(context.Context, *plutus.DiscountCodeValue) (*plutus.DiscountCodeExist, error) {
	panic("unimplemented")
}

// GetActiveDiscountCodes implements a grpc plutus service
func (e *SalesEngine) GetActiveDiscountCodes(context.Context, *plutus.ActiveDiscountsRequest) (*plutus.DiscountCodes, error) {
	panic("unimplemented")
}

// DeleteDiscountCode implements a grpc plutus service
func (e *SalesEngine) DeleteDiscountCode(context.Context, *plutus.DiscountCodeID) (*plutus.DiscountCodes, error) {
	panic("unimplemented")
}
