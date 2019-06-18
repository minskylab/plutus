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
	// error handling omitted
	client := plutus.NewPlutusClient(conn)

	cardToken, err := client.NewCardToken(context.Background(), &plutus.NewCardTokenRequest{
		Card: &plutus.Card{
			Number:  "4111111111111111",
			ExpMont: 9,
			ExpYear: 2020,
			Cvc:     "123",
		},
		Customer: &plutus.Customer{
			// Person: "<PERSON ID>",
			Email: "example@example.com",
		},
	})
	if err != nil {
		panic(err)
	}

	pp.Println(cardToken)
}
