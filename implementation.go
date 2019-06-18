package plutus

import (
	"context"

	plutus "github.com/bregydoc/plutus/proto"
)

// NewCardToken implements a grpc plutus service
func (e *SalesEngine) NewCardToken(c context.Context, p *plutus.NewCardTokenRequest) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// NewCardTokenFromNative implements a grpc plutus service
func (e *SalesEngine) NewCardTokenFromNative(c context.Context, p *plutus.NewCardTokenNativeRequest) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// GetCardTokenOfCustomerByID implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByID(c context.Context, p *plutus.CardTokenByID) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// GetCardTokenOfCustomerByCustomer implements a grpc plutus service
func (e *SalesEngine) GetCardTokenOfCustomerByCustomer(c context.Context, p *plutus.CardTokenByCustomer) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// DeleteCardToken implements a grpc plutus service
func (e *SalesEngine) DeleteCardToken(c context.Context, p *plutus.CardTokenByID) (*plutus.CardToken, error) {
	panic("unimplemented")
}

// NewFastSale implements a grpc plutus service
func (e *SalesEngine) NewFastSale(c context.Context, p *plutus.FastSale) (*plutus.Sale, error) {
	panic("unimplemented")
}

// NewSale implements a grpc plutus service
func (e *SalesEngine) NewSale(c context.Context, p *plutus.NewSaleRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// GetSale implements a grpc plutus service
func (e *SalesEngine) GetSale(c context.Context, p *plutus.SaleIDRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// UpdateSale implements a grpc plutus service
func (e *SalesEngine) UpdateSale(c context.Context, p *plutus.SaleUpdateRequest) (*plutus.Sale, error) {
	panic("unimplemented")
}

// DeliverSaleReceipt implements a grpc plutus service
func (e *SalesEngine) DeliverSaleReceipt(c context.Context, p *plutus.DeliverSaleRequest) (*plutus.DeliverChannelResponse, error) {
	panic("unimplemented")
}

// ChargeSaleByID implements a grpc plutus service
func (e *SalesEngine) ChargeSaleByID(c context.Context, p *plutus.SaleIDRequest) (*plutus.ChargeToken, error) {
	panic("unimplemented")
}

// ChargeSaleWithNativeToken implements a grpc plutus service
func (e *SalesEngine) ChargeSaleWithNativeToken(c context.Context, p *plutus.ChargeWithNativeToken) (*plutus.ChargeToken, error) {
	panic("unimplemented")
}

// DoneSale implements a grpc plutus service
func (e *SalesEngine) DoneSale(c context.Context, p *plutus.SaleIDRequest) (*plutus.Sale, error) {
	panic("unimplemented")
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
