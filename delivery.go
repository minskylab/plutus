package plutus

// DeliveryChannel a delivery channel is a way to represent and send a voucher of yourtransaction.
// examples of delivery channel are: SMTP, SMS, ThermalPrint, etc...
type DeliveryChannel interface {
	SendReceipt(sale *Sale, metadata ...map[string]interface{}) error
}
