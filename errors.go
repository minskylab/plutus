package plutus

import "errors"

var ErrInvalidBridge error = errors.New("your bridge are invalid, please choose a valid bridge")
var ErrNotAvailableBridges error = errors.New("unavailable bridges, please configure one less bridge")
