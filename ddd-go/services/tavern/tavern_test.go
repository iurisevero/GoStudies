package services

import (
	"iurisevero/tavern/domain/product"
	"iurisevero/tavern/services/order"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	custID, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(custID, order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	custID, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(custID, order)
	if err != nil {
		t.Error(err)
	}

}
