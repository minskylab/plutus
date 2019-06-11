package plutus

// Charger is an abstraction layer for your charger ends
type Charger interface {
	NewToken(details CardDetails, kind CardTokenType) (*CardToken, error)
	MakeTransaction(token CardToken)
}
