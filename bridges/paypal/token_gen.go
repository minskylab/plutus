package paypal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (bridge *Bridge) genAccessToken() (*authResponse, error) {
	credentials := fmt.Sprintf("%s:%s", bridge.publicKey, bridge.privateKey)
	crd := base64.StdEncoding.EncodeToString([]byte(credentials))

	body := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest(http.MethodPost, bridge.paypalOAuth, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "api.sandbox.paypal.com")
	req.Header.Add("Authorization", "Basic "+crd)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println("RES CODE:", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		paypalErr := new(paypalError)
		err = json.NewDecoder(res.Body).Decode(paypalErr)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[from paypal] %s", paypalErr.Error)
	}

	token := new(authResponse)
	err = json.NewDecoder(res.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
