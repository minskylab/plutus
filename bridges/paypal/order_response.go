package paypal

import "time"

type orderResponse struct {
	ID            string          `json:"id"`
	Intent        string          `json:"intent"`
	PurchaseUnits []purchaseUnits `json:"purchase_units"`
	Payer         payer           `json:"payer"`
	CreateTime    time.Time       `json:"create_time"`
	UpdateTime    time.Time       `json:"update_time"`
	Links         []links         `json:"links"`
	Status        string          `json:"status"`
}

type amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type payee struct {
	EmailAddress string `json:"email_address"`
	MerchantID   string `json:"merchant_id"`
}

type shippingName struct {
	FullName string `json:"full_name"`
}

type shippingAddress struct {
	AddressLine1 string `json:"address_line_1"`
	AdminArea2   string `json:"admin_area_2"`
	AdminArea1   string `json:"admin_area_1"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}

type shipping struct {
	Name    shippingName    `json:"name"`
	Address shippingAddress `json:"address"`
}

type sellerProtection struct {
	Status            string   `json:"status"`
	DisputeCategories []string `json:"dispute_categories"`
}

type grossAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type paypalFee struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type netAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type sellerReceivableBreakdown struct {
	GrossAmount grossAmount `json:"gross_amount"`
	PaypalFee   paypalFee   `json:"paypal_fee"`
	NetAmount   netAmount   `json:"net_amount"`
}

type links struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type captures struct {
	ID                        string                    `json:"id"`
	Status                    string                    `json:"status"`
	Amount                    amount                    `json:"amount"`
	FinalCapture              bool                      `json:"final_capture"`
	SellerProtection          sellerProtection          `json:"seller_protection"`
	SellerReceivableBreakdown sellerReceivableBreakdown `json:"seller_receivable_breakdown"`
	Links                     []links                   `json:"links"`
	CreateTime                time.Time                 `json:"create_time"`
	UpdateTime                time.Time                 `json:"update_time"`
}

type payments struct {
	Captures []captures `json:"captures"`
}

type purchaseUnits struct {
	ReferenceID string   `json:"reference_id"`
	Amount      amount   `json:"amount"`
	Payee       payee    `json:"payee"`
	Shipping    shipping `json:"shipping"`
	Payments    payments `json:"payments"`
}

type name struct {
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type address struct {
	CountryCode string `json:"country_code"`
}

type payer struct {
	Name         name    `json:"name"`
	EmailAddress string  `json:"email_address"`
	PayerID      string  `json:"payer_id"`
	Address      address `json:"address"`
}
