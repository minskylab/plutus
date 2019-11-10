package main

import (
	"context"
	"encoding/json"
	"os"

	plutus "github.com/bregydoc/plutus/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	pl := plutus.NewPlutusClient(conn)

	card, err := pl.NewCardTokenFromNative(context.Background(), &plutus.NewCardTokenNativeRequest{
		Provider: plutus.Provider_PAYPAL,
		Token:    "08L39595A40835904",
		Type:     plutus.CardTokenType_ONEUSE,
		Customer: &plutus.Customer{
			Name:  "Maria Noriega",
			Email: "maguanorgue1@gmail.com",
		},
	})
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err = enc.Encode(card); err != nil {
		panic(err)
	}

	sale, err := pl.NewFastSale(context.Background(), &plutus.FastSale{
		CustomerEmail: "maguanorgue1@gmail.com",
		Products: []*plutus.Product{
			{Name: "Lima Pass 1", Details: "new kind of pass", Cost: &plutus.Cost{Amount: 20.50, Currency: "USD"}},
		},
	})

	if err != nil {
		panic(err)
	}

	charge, err := pl.ChargeSaleAuto(context.Background(), &plutus.ChargeSaleAutoRequest{
		SaleID:      sale.Id,
		CardTokenID: card.Id,
		Details:     "nothing",
	})

	if err != nil {
		panic(err)
	}

	if err = enc.Encode(charge); err != nil {
		panic(err)
	}

	sale, err = pl.GetSale(context.Background(), &plutus.SaleIDRequest{Id: sale.Id})

	if err = enc.Encode(sale); err != nil {
		panic(err)
	}
}
