package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson06/model"
)

type digitalProductTest struct {
	ID    int
	Name  string
	Price int
	Url   string
}

var digitalProductTests = []digitalProductTest{
	{ID: 1, Name: "坊ちゃん", Price: 2000, Url: "http://a.example.jp"},
	{ID: 2, Name: "スターウォーズ", Price: 4000, Url: "http://b.example.jp"},
}

func TestNewDigitalProduct(t *testing.T) {
	for i := range digitalProductTests {
		test := &digitalProductTests[i]
		c := model.NewDigitalProduct(test.ID, test.Name, test.Price, test.Url)

		if c.ID != test.ID || c.Name != test.Name || c.Price != test.Price || c.Url != test.Url {
			t.Errorf("c.ID, c.Name, c.Price, c.Url = %v, %v, %v, %v want %v, %v, %v, %v",
				c.ID, c.Name, c.Price, c.Url, test.ID, test.Name, test.Price, test.Url)
		}
	}
}
