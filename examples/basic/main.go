package main

import (
	"context"

	"github.com/k0kubun/pp"

	plutus "github.com/bregydoc/plutus/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	pl := plutus.NewPlutusClient(conn)

	card, err := pl.NewCardToken(context.Background(), &plutus.NewCardTokenRequest{
		Card: &plutus.Card{
			Number:  "4111111111111111",
			ExpMont: 9,
			ExpYear: 2020,
			Cvc:     "123",
		},
		Customer: &plutus.Customer{
			Email: "bregymr@gmail.com",
		},
		Type: plutus.CardTokenType_ONEUSE,
	})

	if err != nil {
		panic(err)
	}

	pp.Println(card.Id, card.Value)

	sale, err := pl.NewFastSale(context.Background(), &plutus.FastSale{
		CustomerEmail: "bregymr@gmail.com",
		Products: []*plutus.Product{
			{Name: "Lima Pass 1", Details: "new kind of pass", Cost: &plutus.Cost{Amount: 42.0, Currency: "PEN"}},
		},
	})

	if err != nil {
		panic(err)
	}

	pp.Println(sale.Id, sale.State.String())

	charge, err := pl.ChargeSaleByID(context.Background(), &plutus.ChargeSaleRequest{
		SaleID:      sale.Id,
		CardTokenID: card.Id,
		Details:     "nothing",
	})

	if err != nil {
		panic(err)
	}

	pp.Println(charge.Id, charge.Value)

	sale, err = pl.GetSale(context.Background(), &plutus.SaleIDRequest{Id: sale.Id})

	pp.Println(sale.Id, sale.State.String())

}
