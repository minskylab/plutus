package plutus

// BridgeDescription is used to describe a bridge, it's a fast overview the capabilities of this bridge
type BridgeDescription struct {
	Name                 string
	Version              string
	Type                 ProviderType
	CanGenerateCardToken bool
	CanMakeCharge        bool
	CanMakeRefund        bool
}

// PaymentBridge is an abstraction layer for your charger ends
type PaymentBridge interface {
	Description() *BridgeDescription
	NewToken(details CardDetails, kind CardTokenType) (*CardToken, error)
	MakeCharge(source CardToken, params ChargeParams) (*ChargeToken, error)
	MakeRefund(source ChargeToken, params RefundParams) (*RefundToken, error)
}
