package smtp

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bregydoc/plutus"
	"gopkg.in/mail.v2"
)

// Deliver represents a form of deliver your plutus Sale
type Deliver struct {
	templateFile string
	dialer       *mail.Dialer
	from         string
}

// Config is a wrap for minimal configuration of your smtp deliver way
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// NewSMTPDeliver creates a new devlivery instance
func NewSMTPDeliver(config Config, templateFile string) (*Deliver, error) {
	if config.Host == "" {
		return nil, errors.New("invalid host name")
	}
	port := config.Port
	if port == 0 {
		chunks := strings.Split(config.Host, ":")
		if len(chunks) != 2 {
			return nil, errors.New("invalid port and your host don't have a port in URL")
		}

		var err error
		port, err = strconv.Atoi(chunks[1])
		if err != nil {
			return nil, err
		}
	}

	dialer := mail.NewDialer(config.Host, port, config.Username, config.Password)

	return &Deliver{
		dialer:       dialer,
		templateFile: templateFile,
		from:         config.From,
	}, nil
}

// Name implements a Plutus delivery channel
func (smtp *Deliver) Name() string {
	return "smtp"
}

// SendSale implements a Plutus delivery channel
func (smtp *Deliver) SendSale(from *plutus.Company, sale *plutus.Sale, metadata ...map[string]string) error {
	m := mail.NewMessage()
	if smtp.from != "" {
		m.SetHeader("From", smtp.from)
	} else {
		m.SetHeader("From", from.Support.Email)
	}
	toEmail := ""
	if sale.Customer != nil {
		toEmail = sale.Customer.Email
	}
	if toEmail == "" {
		return errors.New("invalid customer email, please fill your email customer")
	}

	m.SetHeader("To", toEmail)

	meta := map[string]string{}
	if len(metadata) != 0 {
		meta = metadata[0]
	}

	// * If subject exists in metadata
	if subject, ok := meta["subject"]; ok {
		m.SetHeader("Subject", subject)
	} else {
		m.SetHeader("Subject", "Your receipt are ready")
	}

	var temp []byte
	// * If template data is pass throw metadata
	if template, ok := meta["template"]; ok {
		temp = []byte(template)
	} else {
		var err error
		temp, err = ioutil.ReadFile(smtp.templateFile)
		if err != nil {
			return err
		}
	}

	body := bytes.NewBuffer(temp)

	type Data struct {
		Company *plutus.Company
		Sale    *plutus.Sale
	}

	data := Data{
		Company: from,
		Sale:    sale,
	}

	err := template.New("mail_template").Execute(body, data)
	if err != nil {
		return err
	}

	m.SetBody("text/html", body.String())

	return smtp.dialer.DialAndSend(m)
}

// SaleRepresentation implements a Plutus delivery channel
func (smtp *Deliver) SaleRepresentation(from *plutus.Company, sale *plutus.Sale, metadata ...map[string]string) (*plutus.SaleRepresentation, error) {
	meta := map[string]string{}
	if len(metadata) != 0 {
		meta = metadata[0]
	}

	var temp []byte

	// * If template data is pass throw metadata
	if template, ok := meta["template"]; ok {
		temp = []byte(template)
	} else {
		var err error
		temp, err = ioutil.ReadFile(smtp.templateFile)
		if err != nil {
			return nil, err
		}
	}

	body := bytes.NewBuffer(temp)

	type Data struct {
		Company *plutus.Company
		Sale    *plutus.Sale
	}

	data := Data{
		Company: from,
		Sale:    sale,
	}

	err := template.New("mail_template").Execute(body, data)
	if err != nil {
		return nil, err
	}

	return &plutus.SaleRepresentation{
		Data:        body.Bytes(),
		Name:        "invoice",
		Extension:   ".html",
		ContentType: "text/html",
	}, nil

}
