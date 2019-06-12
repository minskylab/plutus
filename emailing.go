package plutus

type EmailDelivery interface {
	SendInvoice(to Customer)
}
