package plutus

// SaleRepresentation is a representation of one sale
type SaleRepresentation struct {
	Data        []byte
	Name        string
	ContentType string
}

// type DeliveryChannel interface {
// 	Name() string
// 	SendSale(from *Company, sale *Sale, metadata ...map[string]string) error
// 	SaleRepresentation(from *Company, sale *Sale, metadata ...map[string]string) (*SaleRepresentation, error)
// }

// DeliveryChannel a delivery channel is a way to represent and send a voucher of yourtransaction.
// examples of delivery channel are: SMTP, SMS, ThermalPrint, etc...
type DeliveryChannel interface {
	Name() string
	DeliverSale(from *Company, sale *Sale, metadata ...map[string]string) (*SaleRepresentation, error)
}
