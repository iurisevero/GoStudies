// Mongo is a mongo implementation of the Customer Repository
package mongo

import (
	"context"
	"iurisevero/tavern/domain/customer"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
	// customer is used to store customers
	customer *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customerAggregate.
// Mongo uses bson so we add tags for that
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromCustomer(c customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a customer.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() customer.Customer {
	c := customer.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}

// Create a new mongodb repository
func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) Update(c customer.Customer) error {
	panic("to implement")
}
