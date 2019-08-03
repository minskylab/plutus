package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bregydoc/plutus"
)

const chargeURL = "https://api.culqi.com/v2/charges"

type chargeResponse struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Outcome struct {
		UserMessage string `json:"user_message"`
	} `json:"outcome"`
}

type antifraudDetails struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty"`
	AddressCity string `json:"address_city,omitempty"`
	PhoneNumber int64  `json:"phone_number,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type chargeParams struct {
	Token            string           `json:"source_id"`
	Email            string           `json:"email,omitempty"`
	CurrencyCode     string           `json:"currency_code,omitempty"`
	Amount           int              `json:"amount,omitempty"`
	Installments     int              `json:"installments,omitempty"`
	Description      string           `json:"description,omitempty"`
	AntifraudDetails antifraudDetails `json:"antifraud_details,omitempty"`
}

type errorResponseFromCulqi struct {
	Object          string `json:"object,omitempty"`
	Type            string `json:"type,omitempty"`
	MerchantMessage string `json:"merchant_message,omitempty"`
}

func (q *PlutusBridge) executeCharge(source plutus.CardToken, params plutus.ChargeParams) (*plutus.ChargeToken, error) {

	apiKey := q.secretKey

	antifraud := antifraudDetails{}
	if params.ExtraInfo != nil {
		chunks := strings.Split(params.ExtraInfo.Name, " ")
		var firstName, lastName string = "", ""

		if len(chunks) >= 2 {
			firstName, lastName = chunks[0], chunks[1]
		}

		antifraud.FirstName = firstName
		antifraud.LastName = lastName
		phone, _ := strconv.Atoi(params.ExtraInfo.Phone)
		antifraud.PhoneNumber = int64(phone)
		if params.ExtraInfo.Location != nil {
			antifraud.Address = params.ExtraInfo.Location.Address
			antifraud.AddressCity = params.ExtraInfo.Location.City
			antifraud.CountryCode = params.ExtraInfo.Location.ZIP
		}
	}

	amount := params.Amount * float64(params.Currency.Multiplier)

	charge := &chargeParams{
		Token:            source.Value,
		Email:            params.Email,
		Description:      params.Details,
		Amount:           int(amount),
		AntifraudDetails: antifraud,
		CurrencyCode:     params.Currency.Name,
	}

	data, err := json.Marshal(charge)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, chargeURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("User-Agent", "plutus/0.2")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		r := new(errorResponseFromCulqi)
		err = json.Unmarshal(data, r)
		if err != nil {
			return nil, err
		}
		if r.Object == "error" {
			return nil, fmt.Errorf("[from culqi] %s", r.MerchantMessage)
		}

	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cResponse := new(chargeResponse)
	err = json.Unmarshal(data, cResponse)
	if err != nil {
		return nil, err
	}

	chargeToken := &plutus.ChargeToken{
		CreatedAt:     time.Now(),
		Value:         cResponse.ID,
		Message:       cResponse.Outcome.UserMessage,
		WithCardToken: source,
	}

	return chargeToken, nil
}
