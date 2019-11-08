package culqi

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const baseCheckoutURL = "https://checkout.culqi.com/"
const sessionHeaderName = "X-CULQI-SESSIONID"

// CheckoutParams ...
type CheckoutParams struct {
	Title        string
	Currency     string
	Description  string
	Amount       string
	Logo         string
	Installments bool
	Orders       string
}

func encodeString(value string) string {
	return base64.URLEncoding.EncodeToString([]byte(value))
}

func (bridge *Bridge) generateCheckoutURL(params ...CheckoutParams) (string, error) {
	ch, err := url.Parse(baseCheckoutURL)
	if err != nil {
		return "", err
	}
	v, err := url.ParseQuery("")
	if err != nil {
		return "", err
	}

	p := defaultCheckout
	if len(params) > 0 {
		p = params[0]
	}

	v.Add("public_key", bridge.publicKey)
	v.Add("title", encodeString(p.Title))
	v.Add("currency", encodeString(p.Currency))
	v.Add("description", encodeString(p.Description))
	v.Add("amount", encodeString(p.Amount))
	v.Add("logo", encodeString(p.Logo))
	v.Add("installments", strconv.FormatBool(p.Installments))
	v.Add("orders", encodeString(p.Orders))

	ch.RawQuery = v.Encode()
	return ch.String(), nil
}

func (bridge *Bridge) getNewSessionID(params ...CheckoutParams) (string, error) {
	uri, err := bridge.generateCheckoutURL(params...)
	if err != nil {
		return "", err
	}

	r, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	for _, c := range r.Cookies() {
		log.Println(c.String())
	}
	if len(r.Cookies()) == 0 {
		return "", ErrCulqiNotWorking
	}

	sess := ""
	for _, c := range r.Cookies() {
		if c.Name == sessionHeaderName {
			if len(c.Value) > len(sess) {
				sess = c.Value
			}
		}

	}

	return sess, nil
}
