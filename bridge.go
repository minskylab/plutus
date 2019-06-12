package plutus

// PaymentBridge is an abstraction layer for your charger ends
type PaymentBridge interface {
	NewToken(details CardDetails, kind CardTokenType) (*CardToken, error)
	MakeCharge(source CardToken, params ChargeParams) (*ChargeToken, error)
	MakeRefund(source ChargeToken, params RefundParams) (*RefundToken, error)
}
