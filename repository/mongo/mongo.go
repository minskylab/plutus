package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "plutus"

const customerCollection = "customers"
const cardCollection = "cards"
const chargeCollection = "charges"
const saleCollection = "sales"

type Repository struct {
	URI       string
	client    *mongo.Client
	db        *mongo.Database
	customers *mongo.Collection
	cards     *mongo.Collection
	charges   *mongo.Collection
	sales     *mongo.Collection
}

func NewRepository(c context.Context, URI string) (*Repository, error) {
	client, err := mongo.Connect(c, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	customers := db.Collection(customerCollection)
	cards := db.Collection(cardCollection)
	charges := db.Collection(chargeCollection)
	sales := db.Collection(saleCollection)

	uniqueOption := options.Index().SetUnique(true)

	if _, err = customers.Indexes().CreateMany(c, []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"ID": 1}, Options: uniqueOption},
		mongo.IndexModel{Keys: bson.M{"Email": 1}, Options: uniqueOption},
	}); err != nil {
		return nil, err
	}

	if _, err = cards.Indexes().CreateMany(c, []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"ID": 1}, Options: uniqueOption},
		mongo.IndexModel{Keys: bson.M{"Value": 1}, Options: uniqueOption},
	}); err != nil {
		return nil, err
	}

	if _, err = charges.Indexes().CreateMany(c, []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"ID": 1}, Options: uniqueOption},
		mongo.IndexModel{Keys: bson.M{"Value": 1}, Options: uniqueOption},
	}); err != nil {
		return nil, err
	}

	if _, err = sales.Indexes().CreateMany(c, []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"ID": 1}, Options: uniqueOption},
	}); err != nil {
		return nil, err
	}

	repo := &Repository{
		URI:       URI,
		client:    client,
		db:        db,
		customers: customers,
		cards:     cards,
		charges:   charges,
		sales:     sales,
	}

	return repo, nil
}
