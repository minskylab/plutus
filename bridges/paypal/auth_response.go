package paypal

import "time"

type authResponse struct {
	Scope       string    `json:"scope"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	AppID       string    `json:"app_id"`
	ExpiresIn   int       `json:"expires_in"`
	Nonce       time.Time `json:"nonce"`
}
