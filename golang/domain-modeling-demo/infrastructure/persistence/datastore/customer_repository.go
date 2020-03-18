package datastore

import (
	"github.com/tabakazu/domain-modeling-demo/domain/model/collection"
	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/repository"
)

type CustomerRepository struct {
	*Config
}

func NewCustomerRepository() repository.CustomerRepository {
	config := NewConfig()
	return &CustomerRepository{Config: config}
}

func (r *CustomerRepository) Save(customer entity.Customer) (entity.Customer, error) {
	if err := r.Conn.Create(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *CustomerRepository) FindAll() (collection.CustomerCollection, error) {
	var customerCollection collection.CustomerCollection
	if err := r.Conn.Find(&customerCollection.List).Error; err != nil {
		return customerCollection, err
	}
	return customerCollection, nil
}
