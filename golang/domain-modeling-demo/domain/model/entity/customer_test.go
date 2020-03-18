package entity_test

import (
	"testing"

	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
)

func Test_NewCustomer(t *testing.T) {
	name := valueobject.CustomerName{Value: "test customer"}
	address := valueobject.Address{Street: "Nihon-ku Minato-shi 1-1-1", City: "Tokyo-to", Country: "Japan"}
	customer := entity.NewCustomer(name, address)

	if customer == nil {
		t.Errorf("customer = %v want %v", customer, nil)
	}
	if customer.Name.Value != name.Value {
		t.Errorf("customer.Name.Value = %v want %v", customer.Name.Value, name.Value)
	}
	if customer.Address != address {
		t.Errorf("customer.Address = %v want %v", customer.Address, address)
	}
}
