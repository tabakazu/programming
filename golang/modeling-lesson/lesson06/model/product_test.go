package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson06/model"
)

type productTest struct {
	ID    int
	Name  string
	Price int
}

var productTests = []productTest{
	{ID: 1, Name: "坊ちゃん", Price: 2000},
	{ID: 2, Name: "スターウォーズ", Price: 4000},
}

func TestNewProduct(t *testing.T) {
	for i := range productTests {
		test := &productTests[i]
		c := model.NewProduct(test.ID, test.Name, test.Price)

		if c.ID != test.ID || c.Name != test.Name || c.Price != test.Price {
			t.Errorf("c.ID, c.Name, c.Price = %v, %v, %v want %v, %v, %v", c.ID, c.Name, c.Price, test.ID, test.Name, test.Price)
		}
	}
}
