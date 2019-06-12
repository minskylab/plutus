package culqi

import (
	"errors"
	"strconv"
	"time"

	"github.com/bregydoc/plutus"
)

func validateNumber(number string) bool {
	var sum int
	var alternate bool

	numberLen := len(number)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(number[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}

func validateCardDetails(card plutus.CardDetails) error {
	year := card.Expiration.Year

	month := card.Expiration.Month

	if month < 1 || 12 < month {
		return errors.New("invalid month")
	}

	if year < time.Now().UTC().Year() {
		return errors.New("credit card has expired")
	}

	if year == time.Now().UTC().Year() && month < int(time.Now().UTC().Month()) {
		return errors.New("credit card has expired")
	}

	if len(card.CVV) < 3 || len(card.CVV) > 4 {
		return errors.New("invalid CVV")
	}

	if len(card.Number) < 13 {
		return errors.New("invalid credit card number")
	}

	if valid := validateNumber(card.Number); !valid {
		return errors.New("invalid credit card number")
	}

	return nil
}
