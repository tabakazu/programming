package repository

import (
	"github.com/tabakazu/domain-modeling-demo/domain/model/collection"
	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
)

type CustomerRepository interface {
	Save(post entity.Customer) (entity.Customer, error)
	FindAll() (collection.CustomerCollection, error)
}
