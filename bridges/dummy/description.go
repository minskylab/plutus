package dummy

import "github.com/bregydoc/plutus"

var dummyDesc = &plutus.BridgeDescription{
	Name:                 "dummy",
	Version:              "0.1.0",
	CanGenerateCardToken: true,
	CanMakeCharge:        true,
	CanMakeRefund:        true,
}