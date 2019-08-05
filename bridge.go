package plutus

// BridgeDescription is used to describe a bridge, is a fast overview of the capabilities of this bridge
type BridgeDescription struct {
	Name                 string
	Version              string
	CanGenerateCardToken bool
	CanMakeCharge        bool
	CanMakeRefund        bool
}

// PaymentBridge is an abstraction layer for your charger ends
type PaymentBridge interface {
	Describe() *BridgeDescription
	NewToken(details CardDetails, kind CardTokenType) (*CardToken, error)
	MakeCharge(source CardToken, params ChargeParams) (*ChargeToken, error)
	MakeRefund(source ChargeToken, params RefundParams) (*RefundToken, error)
}
