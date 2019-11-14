package smtp

import (
	"bytes"
	"errors"
	"html/template"
	"strconv"
	"strings"

	"github.com/bregydoc/plutus"
	"gopkg.in/mail.v2"
)

// Deliver represents a form of deliver your plutus Sale
type Deliver struct {
	templateEngine TemplateEngine
	dialer         *mail.Dialer
	from           string
}

// Config is a wrap for minimal configuration of your smtp deliver way
type Config struct {
	Host     string
	Port     int64
	Username string
	Password string
	From     string
}

// NewSMTPDeliver creates a new devlivery instance
func NewDeliver(config Config, engine TemplateEngine) (*Deliver, error) {
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
		p, err := strconv.Atoi(chunks[1])
		if err != nil {
			return nil, err
		}
		port = int64(p)
	}

	dialer := mail.NewDialer(config.Host, int(port), config.Username, config.Password)

	return &Deliver{
		dialer:         dialer,
		templateEngine: engine,
		from:           config.From,
	}, nil
}

// Name implements a Plutus delivery channel
func (smtp *Deliver) Name() string {
	return "smtp"
}

// DeliverSale implements a Plutus delivery channel
func (smtp *Deliver) DeliverSale(from *plutus.Company, sale *plutus.Sale, metadata ...map[string]string) (*plutus.SaleRepresentation, error) {
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
		return nil, errors.New("invalid customer email, please fill your email customer")
	}

	m.SetHeader("To", toEmail)

	meta := map[string]string{}
	if len(metadata) != 0 {
		for _, m := range metadata {
			for k, e := range m {
				meta[k] = e
			}
		}
	}

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
		t, err := smtp.templateEngine.GetTemplate()
		if err != nil {
			return nil, err
		}
		temp = []byte(t)
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

	m.SetBody("text/html", body.String())

	if err = smtp.dialer.DialAndSend(m); err != nil {
		return nil, err
	}

	repr := &plutus.SaleRepresentation{
		Data:        body.Bytes(),
		Name:        "invoice",
		ContentType: "text/html",
	}

	return repr, nil
}
