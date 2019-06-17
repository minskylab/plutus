package repository

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/bregydoc/plutus"
)

// BoltRepository is an repository implementation of plutus repository
type BoltRepository struct {
	db *storm.DB
}

// NewBoltRepository returns a new instance of a plutus repositry implementation
func NewBoltRepository(path string) (*BoltRepository, error) {
	db, err := storm.Open(path)
	if err != nil {
		return nil, err
	}

	return &BoltRepository{
		db: db,
	}, nil
}

// SaveCustomer implements a plutus repository
func (repo *BoltRepository) SaveCustomer(customer *plutus.Customer) (*plutus.Customer, error) {
	if customer.ID == "" {
		customer.FillID()
	}
	err := repo.db.From("customers").Save(customer)
	if err != nil {
		return nil, err
	}
	return repo.GetCustomer(customer.ID)
}

// GetCustomer implements a plutus repository
func (repo *BoltRepository) GetCustomer(ID string) (*plutus.Customer, error) {
	customer := new(plutus.Customer)
	err := repo.db.From("customers").One("ID", ID, customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// UpdateCustomer implements a plutus repository
func (repo *BoltRepository) UpdateCustomer(ID string, updatePayload plutus.Customer) (*plutus.Customer, error) {
	customer, err := repo.GetCustomer(ID)
	if err != nil {
		return nil, err
	}

	if updatePayload.Email != "" {
		customer.Email = updatePayload.Email
	}

	if updatePayload.Name != "" {
		customer.Name = updatePayload.Name
	}

	if updatePayload.Person != "" {
		customer.Person = updatePayload.Person
	}

	if updatePayload.Phone != "" {
		customer.Phone = updatePayload.Phone
	}

	if updatePayload.Location != nil {
		if updatePayload.Location.Address != "" {
			customer.Location.Address = updatePayload.Location.Address
		}
		if updatePayload.Location.City != "" {
			customer.Location.City = updatePayload.Location.City
		}
		if updatePayload.Location.State != "" {
			customer.Location.State = updatePayload.Location.State
		}
		if updatePayload.Location.ZIP != "" {
			customer.Location.ZIP = updatePayload.Location.ZIP
		}

	}

	err = repo.db.From("customers").Update(customer)
	if err != nil {
		return nil, err
	}

	return customer, err
}

// RemoveCustomer implements a plutus repository
func (repo *BoltRepository) RemoveCustomer(ID string) (*plutus.Customer, error) {
	customer, err := repo.GetCustomer(ID)
	if err != nil {
		return nil, err
	}

	err = repo.db.From("customers").DeleteStruct(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// SaveCardToken implements a plutus repository
func (repo *BoltRepository) SaveCardToken(cardToken *plutus.CardToken) (*plutus.CardToken, error) {
	if cardToken.ID == "" {
		cardToken.FillID()
	}

	err := repo.db.From("cards").Save(cardToken)
	if err != nil {
		return nil, err
	}

	return repo.GetCardToken(cardToken.ID)

}

// GetCardToken implements a plutus repository
func (repo *BoltRepository) GetCardToken(ID string) (*plutus.CardToken, error) {
	card := new(plutus.CardToken)
	err := repo.db.From("cards").One("ID", ID, card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

// UpdateCardToken implements a plutus repository
func (repo *BoltRepository) UpdateCardToken(ID string, updatePayload plutus.CardToken) (*plutus.CardToken, error) {
	card, err := repo.GetCardToken(ID)
	if err != nil {
		return nil, err
	}

	if updatePayload.Type != "" {
		card.Type = updatePayload.Type
	}

	if updatePayload.Value != "" {
		card.Value = updatePayload.Value
	}

	if updatePayload.WithCard.Number != "" {
		card.WithCard.Number = updatePayload.WithCard.Number
	}

	if updatePayload.WithCard.ExpirationYear != 0 {
		card.WithCard.ExpirationYear = updatePayload.WithCard.ExpirationYear
	}

	if updatePayload.WithCard.Customer != nil {
		card.WithCard.Customer = updatePayload.WithCard.Customer
	}

	err = repo.db.From("cards").Update(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// RemoveCardToken implements a plutus repository
func (repo *BoltRepository) RemoveCardToken(ID string) (*plutus.CardToken, error) {
	card, err := repo.GetCardToken(ID)
	if err != nil {
		return nil, err
	}

	err = repo.db.From("cards").DeleteStruct(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// SaveChargeToken implements a plutus repository
func (repo *BoltRepository) SaveChargeToken(chargeToken *plutus.ChargeToken) (*plutus.ChargeToken, error) {
	if chargeToken.ID == "" {
		chargeToken.FillID()
	}

	err := repo.db.From("charges").Save(chargeToken)
	if err != nil {
		return nil, err
	}

	return repo.GetChargeToken(chargeToken.ID)

}

// GetChargeToken implements a plutus repository
func (repo *BoltRepository) GetChargeToken(ID string) (*plutus.ChargeToken, error) {
	charge := new(plutus.ChargeToken)
	err := repo.db.From("charges").One("ID", ID, charge)
	if err != nil {
		return nil, err
	}
	return charge, nil
}

// UpdateChargeToken implements a plutus repository
func (repo *BoltRepository) UpdateChargeToken(ID string, updatePayload plutus.ChargeToken) (*plutus.ChargeToken, error) {
	charge, err := repo.GetChargeToken(ID)
	if err != nil {
		return nil, err
	}

	fmt.Println("[WARNING] You can not to update a charge")
	err = repo.db.From("charges").Update(charge)
	if err != nil {
		return nil, err
	}

	return charge, nil
}

// RemoveChargeToken implements a plutus repository
func (repo *BoltRepository) RemoveChargeToken(ID string) (*plutus.ChargeToken, error) {
	charge, err := repo.GetChargeToken(ID)
	if err != nil {
		return nil, err
	}

	err = repo.db.From("charges").DeleteStruct(charge)
	if err != nil {
		return nil, err
	}

	return charge, nil
}

// SaveSale implements a plutus repository
func (repo *BoltRepository) SaveSale(sale *plutus.Sale) (*plutus.Sale, error) {
	if sale.ID == "" {
		sale.FillID()
	}

	err := repo.db.From("sales").Save(sale)
	if err != nil {
		return nil, err
	}

	return repo.GetSale(sale.ID)

}

// GetSale implements a plutus repository
func (repo *BoltRepository) GetSale(ID string) (*plutus.Sale, error) {
	sale := new(plutus.Sale)
	err := repo.db.From("sales").One("ID", ID, sale)
	if err != nil {
		return nil, err
	}
	return sale, nil
}

// UpdateSale implements a plutus repository
func (repo *BoltRepository) UpdateSale(ID string, updatePayload plutus.Sale) (*plutus.Sale, error) {
	panic("unimplemented")
}

// RemoveSale implements a plutus repository
func (repo *BoltRepository) RemoveSale(ID string) (*plutus.Sale, error) {
	sale, err := repo.GetSale(ID)
	if err != nil {
		return nil, err
	}

	err = repo.db.From("sales").DeleteStruct(sale)
	if err != nil {
		return nil, err
	}

	return sale, nil
}
