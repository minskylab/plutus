package plutus

import (
	plutus "github.com/bregydoc/plutus/proto"
	"github.com/golang/protobuf/ptypes"
)

func locationToProto(location Location) *plutus.Location {
	return &plutus.Location{
		Address: location.Address,
		City:    location.City,
		State:   location.State,
		Zip:     location.ZIP,
	}
}

func locationFromProto(location *plutus.Location) Location {
	return Location{
		Address: location.Address,
		City:    location.City,
		State:   location.State,
		ZIP:     location.Zip,
	}
}

func customerToProto(customer Customer) *plutus.Customer {
	return &plutus.Customer{
		Id:       customer.ID,
		Email:    customer.Email,
		Name:     customer.Name,
		Phone:    customer.Phone,
		Location: locationToProto(*customer.Location),
	}
}

func customerFromProto(customer *plutus.Customer) Customer {
	loc := locationFromProto(customer.Location)
	return Customer{
		ID:       customer.Id,
		Email:    customer.Email,
		Name:     customer.Name,
		Phone:    customer.Phone,
		Location: &loc,
	}
}

func encodedCardDetailsToProto(details EncodedCardDetails) *plutus.EncodedCardDetails {
	return &plutus.EncodedCardDetails{
		Number:         details.Number,
		ExpirationYear: int32(details.ExpirationYear),
		Customer:       customerToProto(*details.Customer),
	}
}

func encodedCardDetailsFromProto(details *plutus.EncodedCardDetails) EncodedCardDetails {
	customer := customerFromProto(details.Customer)
	return EncodedCardDetails{
		Number:         details.Number,
		ExpirationYear: int(details.ExpirationYear),
		Customer:       &customer,
	}
}

func cardTokenToProto(token CardToken) *plutus.CardToken {
	createdAt, _ := ptypes.TimestampProto(token.CreatedAt)
	t := plutus.CardTokenType_ONEUSE
	if token.Type == RecurrentToken {
		t = plutus.CardTokenType_RECURRENT
	}
	return &plutus.CardToken{
		CreatedAt: createdAt,
		Id:        token.ID,
		Type:      t,
		Value:     token.Value,
		WithCard:  encodedCardDetailsToProto(token.WithCard),
	}
}

func cardTokenFromProto(token *plutus.CardToken) CardToken {
	createdAt, _ := ptypes.Timestamp(token.CreatedAt)
	t := OneUseToken
	if token.Type == plutus.CardTokenType_RECURRENT {
		t = RecurrentToken
	}
	return CardToken{
		CreatedAt: createdAt,
		ID:        token.Id,
		Type:      t,
		Value:     token.Value,
		WithCard:  encodedCardDetailsFromProto(token.WithCard),
	}
}

func chargeTokenToProto(token ChargeToken) *plutus.ChargeToken {
	createdAt, _ := ptypes.TimestampProto(token.CreatedAt)
	return &plutus.ChargeToken{
		CreatedAt:     createdAt,
		Id:            token.ID,
		Message:       token.Message,
		Value:         token.Value,
		WithCardToken: cardTokenToProto(token.WithCardToken),
	}
}

func chargeTokenFromProto(token *plutus.ChargeToken) ChargeToken {
	createdAt, _ := ptypes.Timestamp(token.CreatedAt)
	return ChargeToken{
		CreatedAt:     createdAt,
		ID:            token.Id,
		Message:       token.Message,
		Value:         token.Value,
		WithCardToken: cardTokenFromProto(token.WithCardToken),
	}
}

func costToProto(cost Cost) *plutus.Cost {
	return &plutus.Cost{
		Amount:   cost.Amount,
		Currency: cost.Currency.Name,
	}
}

func costFromProto(cost *plutus.Cost) Cost {
	//! Check that later
	curr := AvailableCurrencies[cost.Currency]
	return Cost{
		Amount:   cost.Amount,
		Currency: curr,
	}
}

func productToProto(product Product) *plutus.Product {
	return &plutus.Product{
		Name:    product.Name,
		Details: product.Details,
		Cost:    costToProto(product.Cost),
	}
}

func productFromProto(product *plutus.Product) Product {
	return Product{
		Name:    product.Name,
		Details: product.Details,
		Cost:    costFromProto(product.Cost),
	}
}

func productsToProto(products []Product) []*plutus.Product {
	pProducts := make([]*plutus.Product, 0)
	for _, p := range products {
		pProducts = append(pProducts, productToProto(p))
	}
	return pProducts
}

func productsFromProto(products []*plutus.Product) []Product {
	nProducts := make([]Product, 0)
	for _, p := range products {
		nProducts = append(nProducts, productFromProto(p))
	}
	return nProducts
}

func saleStateToProto(state SaleState) plutus.SaleState {
	switch state {
	case Null:
		// Draft is the earliest state of sale
		return plutus.SaleState_NULL
	case Draft:
		// Signed is a signed and final modeling sale
		return plutus.SaleState_DRAFT
	case Signed:
		// PaidOut is a charged sale
		return plutus.SaleState_SIGNED
	case PaidOut:
		// Done is a done sale
		return plutus.SaleState_PAIDOUT
	case Done:
		return plutus.SaleState_DONE
	}
	return plutus.SaleState_NULL
}

func saleStateFromProto(state plutus.SaleState) SaleState {
	switch state {
	case plutus.SaleState_NULL:
		// Draft is the earliest state of sale
		return Null
	case plutus.SaleState_DRAFT:
		// Signed is a signed and final modeling sale
		return Draft
	case plutus.SaleState_SIGNED:
		// PaidOut is a charged sale
		return Signed
	case plutus.SaleState_PAIDOUT:
		// Done is a done sale
		return PaidOut
	case plutus.SaleState_DONE:
		return Done
	}
	return Null
}

func discountUseRecordToProto(record DiscountUseRecord) *plutus.DiscountUseRecord {
	at, _ := ptypes.TimestampProto(record.At)
	return &plutus.DiscountUseRecord{
		At: at,
		By: customerToProto(*record.By),
	}
}

func discountUseRecordFromProto(record *plutus.DiscountUseRecord) DiscountUseRecord {
	customer := customerFromProto(record.By)
	at, _ := ptypes.Timestamp(record.At)
	return DiscountUseRecord{
		At: at,
		By: &customer,
	}
}

func discountUsesRecordToProto(records []DiscountUseRecord) []*plutus.DiscountUseRecord {
	pRecords := make([]*plutus.DiscountUseRecord, 0)
	for _, r := range records {
		pRecords = append(pRecords, discountUseRecordToProto(r))
	}
	return pRecords
}

func discountUsesRecordFromProto(records []*plutus.DiscountUseRecord) []DiscountUseRecord {
	nRecords := make([]DiscountUseRecord, 0)
	for _, r := range records {
		nRecords = append(nRecords, discountUseRecordFromProto(r))
	}
	return nRecords
}

func discountToProto(discount Discount) *plutus.Discount {
	discountType := plutus.DiscountType_PERCENT
	if discount.Type == StaticDiscount {
		discountType = plutus.DiscountType_STATIC
	}
	return &plutus.Discount{
		Amount:   discount.Amount,
		Currency: discount.Currency.Name,
		Percent:  discount.Percent,
		Type:     discountType,
	}
}

func discountFromProto(discount *plutus.Discount) Discount {
	discountType := PercentDiscount
	if discount.Type == plutus.DiscountType_STATIC {
		discountType = StaticDiscount
	}
	curr := AvailableCurrencies[discount.Currency]
	return Discount{
		Amount:   discount.Amount,
		Currency: curr,
		Percent:  discount.Percent,
		Type:     discountType,
	}
}

func discountCodeToProto(code DiscountCode) *plutus.DiscountCode {
	start, _ := ptypes.TimestampProto(code.Start)
	end, _ := ptypes.TimestampProto(code.End)
	return &plutus.DiscountCode{
		Id:      code.ID,
		Code:    code.Code,
		MaxUses: int32(code.MaxUses),
		Start:   start,
		End:     end,
		Uses:    discountUsesRecordToProto(code.Uses),
		Value:   discountToProto(*code.Value),
	}
}

func discountCodeFromProto(code *plutus.DiscountCode) DiscountCode {
	start, _ := ptypes.Timestamp(code.Start)
	end, _ := ptypes.Timestamp(code.End)
	val := discountFromProto(code.Value)
	return DiscountCode{
		ID:      code.Id,
		Code:    code.Code,
		MaxUses: int(code.MaxUses),
		Start:   start,
		End:     end,
		Uses:    discountUsesRecordFromProto(code.Uses),
		Value:   &val,
	}
}

func discountCodesToProto(codes []DiscountCode) []*plutus.DiscountCode {
	pCodes := make([]*plutus.DiscountCode, 0)
	for _, d := range codes {
		pCodes = append(pCodes, discountCodeToProto(d))
	}
	return pCodes
}

func discountCodesFromProto(codes []*plutus.DiscountCode) []DiscountCode {
	nCodes := make([]DiscountCode, 0)
	for _, d := range codes {
		nCodes = append(nCodes, discountCodeFromProto(d))
	}
	return nCodes
}

func saleToProto(sale Sale) *plutus.Sale {
	createdAt, _ := ptypes.TimestampProto(sale.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(sale.UpdatedAt)

	return &plutus.Sale{
		Id:            sale.ID,
		CardToken:     cardTokenToProto(*sale.CardToken),
		Charge:        chargeTokenToProto(*sale.Charge),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Customer:      customerToProto(*sale.Customer),
		CurrencyToPay: sale.CurrencyToPay.Name,
		Products:      productsToProto(sale.Products),
		State:         saleStateToProto(sale.State),
		DiscountCodes: discountCodesToProto(sale.DiscountCodes),
	}
}

func saleFromProto(sale *plutus.Sale) Sale {
	createdAt, _ := ptypes.Timestamp(sale.CreatedAt)
	updatedAt, _ := ptypes.Timestamp(sale.UpdatedAt)

	card := cardTokenFromProto(sale.CardToken)
	charge := chargeTokenFromProto(sale.Charge)
	customer := customerFromProto(sale.Customer)
	curr := AvailableCurrencies[sale.CurrencyToPay]
	return Sale{
		ID:            sale.Id,
		CardToken:     &card,
		Charge:        &charge,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Customer:      &customer,
		CurrencyToPay: curr,
		Products:      productsFromProto(sale.Products),
		State:         saleStateFromProto(sale.State),
		DiscountCodes: discountCodesFromProto(sale.DiscountCodes),
	}
}