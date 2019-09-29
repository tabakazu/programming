package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson06/model"
)

type realProductTest struct {
	ID         int
	Name       string
	Price      int
	StockCount int
}

var realProductTests = []realProductTest{
	{ID: 1, Name: "坊ちゃん", Price: 2000, StockCount: 5},
	{ID: 2, Name: "スターウォーズ", Price: 4000, StockCount: 1},
}

func TestNewRealProduct(t *testing.T) {
	for i := range realProductTests {
		test := &realProductTests[i]
		c := model.NewRealProduct(test.ID, test.Name, test.Price, test.StockCount)

		if c.ID != test.ID || c.Name != test.Name || c.Price != test.Price || c.StockCount != test.StockCount {
			t.Errorf("c.ID, c.Name, c.Price, c.StockCount = %v, %v, %v, %v want %v, %v, %v, %v",
				c.ID, c.Name, c.Price, c.StockCount, test.ID, test.Name, test.Price, test.StockCount)
		}
	}
}
