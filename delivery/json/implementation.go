package json

import (
	"encoding/json"
	"io"

	"github.com/bregydoc/plutus"
)

type Deliver struct {
	encoder *json.Encoder
}

func NewDeliver(w io.Writer) (*Deliver, error) {
	return &Deliver{encoder: json.NewEncoder(w)}, nil
}

func (d *Deliver) Name() string {
	return "json"
}

func (d *Deliver) DeliverSale(from *plutus.Company, sale *plutus.Sale, metadata ...map[string]string) (*plutus.SaleRepresentation, error) {
	type Data struct {
		Company *plutus.Company
		Sale    *plutus.Sale
	}

	data := Data{
		Company: from,
		Sale:    sale,
	}

	if err := d.encoder.Encode(data); err != nil {
		return nil, err
	}

	dataSerialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	repr := &plutus.SaleRepresentation{
		Data:        dataSerialized,
		Name:        "invoice",
		ContentType: "application/json",
	}

	return repr, nil
}
