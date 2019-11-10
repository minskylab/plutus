package mongo

import (
	"context"

	"github.com/bregydoc/plutus"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) SaveCustomer(c context.Context, customer *plutus.Customer) (*plutus.Customer, error) {
	if customer.ID == "" {
		customer.FillID()
	}
	if _, err := repo.customers.InsertOne(c, customer); err != nil {
		return nil, err
	}

	return repo.GetCustomer(c, customer.ID)
}

func (repo *Repository) GetCustomer(c context.Context, ID string) (*plutus.Customer, error) {
	res := repo.customers.FindOne(c, bson.M{"ID": ID})

	if res.Err() != nil {
		return nil, res.Err()
	}

	customer := new(plutus.Customer)
	if err := res.Decode(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo *Repository) UpdateCustomer(c context.Context, ID string, updatePayload plutus.Customer) (*plutus.Customer, error) {
	// repo.customers.UpdateOne(c, )
	// TODO: Implement update with $set
	panic("unimplemented")
}

func (repo *Repository) RemoveCustomer(c context.Context, ID string) (*plutus.Customer, error) {
	customer, err := repo.GetCustomer(c, ID)
	if err != nil {
		return nil, err
	}

	if _, err := repo.customers.DeleteOne(c, bson.M{"ID": ID}); err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo *Repository) SaveCardToken(c context.Context, cardToken *plutus.CardToken) (*plutus.CardToken, error) {
	if cardToken.ID == "" {
		cardToken.FillID()
	}

	if _, err := repo.cards.InsertOne(c, cardToken); err != nil {
		return nil, err
	}

	return repo.GetCardToken(c, cardToken.ID)
}

func (repo *Repository) GetCardToken(c context.Context, ID string) (*plutus.CardToken, error) {
	res := repo.cards.FindOne(c, bson.M{"ID": ID})
	if res.Err() != nil {
		return nil, res.Err()
	}
	card := new(plutus.CardToken)

	if err := res.Decode(card); err != nil {
		return nil, err
	}

	return card, nil
}

func (repo *Repository) UpdateCardToken(c context.Context, ID string, updatePayload plutus.CardToken) (*plutus.CardToken, error) {
	panic("unimplemented")
	// TODO: implement update with $set and validating the upload payload
}

func (repo *Repository) RemoveCardToken(c context.Context, ID string) (*plutus.CardToken, error) {
	card, err := repo.GetCardToken(c, ID)
	if err != nil {
		return nil, err
	}

	if _, err := repo.cards.DeleteOne(c, bson.M{"ID": ID}); err != nil {
		return nil, err
	}
	return card, nil
}

func (repo *Repository) SaveChargeToken(c context.Context, chargeToken *plutus.ChargeToken) (*plutus.ChargeToken, error) {
	if chargeToken.ID == "" {
		chargeToken.FillID()
	}

	if _, err := repo.cards.InsertOne(c, chargeToken); err != nil {
		return nil, err
	}

	return repo.GetChargeToken(c, chargeToken.ID)
}

func (repo *Repository) GetChargeToken(c context.Context, ID string) (*plutus.ChargeToken, error) {
	res := repo.charges.FindOne(c, bson.M{"ID": ID})

	if res.Err() != nil {
		return nil, res.Err()
	}

	charge := new(plutus.ChargeToken)
	if err := res.Decode(charge); err != nil {
		return nil, err
	}

	return charge, nil
}

func (repo *Repository) UpdateChargeToken(c context.Context, ID string, updatePayload plutus.ChargeToken) (*plutus.ChargeToken, error) {
	panic("unimplemented")
	// TODO: implement update with $set and validating the upload payload
}

func (repo *Repository) RemoveChargeToken(c context.Context, ID string) (*plutus.ChargeToken, error) {
	charge, err := repo.GetChargeToken(c, ID)
	if err != nil {
		return nil, err
	}

	if _, err := repo.charges.DeleteOne(c, bson.M{"ID": ID}); err != nil {
		return nil, err
	}
	return charge, nil
}

func (repo *Repository) SaveSale(c context.Context, sale *plutus.Sale) (*plutus.Sale, error) {
	if sale.ID == "" {
		sale.FillID()
	}
	if _, err := repo.sales.InsertOne(c, sale); err != nil {
		return nil, err
	}

	return repo.GetSale(c, sale.ID)
}

func (repo *Repository) GetSale(c context.Context, ID string) (*plutus.Sale, error) {
	res := repo.sales.FindOne(c, bson.M{"ID": ID})

	if res.Err() != nil {
		return nil, res.Err()
	}

	sale := new(plutus.Sale)
	if err := res.Decode(sale); err != nil {
		return nil, err
	}

	return sale, nil

}

func (repo *Repository) UpdateSale(c context.Context, ID string, updatePayload plutus.Sale) (*plutus.Sale, error) {
	panic("unimplemented")
	// TODO: implement update with $set and validating the upload payload
}

func (repo *Repository) RemoveSale(c context.Context, ID string) (*plutus.Sale, error) {
	sale, err := repo.GetSale(c, ID)
	if err != nil {
		return nil, err
	}

	if _, err := repo.sales.DeleteOne(c, bson.M{"ID": ID}); err != nil {
		return nil, err
	}
	return sale, nil
}
