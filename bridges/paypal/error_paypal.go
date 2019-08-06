package paypal

type paypalError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
