package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson05/model"
)

type businessCustomerTest struct {
	ID      int
	Name    string
	Email   string
	Capital int
}

var (
	businessCustomerTests = []businessCustomerTest{
		{ID: 1, Name: "鈴木テクノロジーズ", Email: "suzuki@example.jp", Capital: 1000000},
		{ID: 2, Name: "佐藤商社", Email: "sato@example.jp", Capital: 20000000},
	}
)

func TestNewBusinessCustomer(t *testing.T) {
	for i := range businessCustomerTests {
		test := &businessCustomerTests[i]
		c := model.NewBusinessCustomer(test.ID, test.Name, test.Email, test.Capital)

		if c.ID != test.ID || c.Name != test.Name || c.Email != test.Email || c.Capital != test.Capital {
			t.Errorf("c.ID, c.Name, c.Email, c.Capital = %v, %v, %v, %v want %v, %v, %v, %v",
				c.ID, c.Name, c.Email, c.Capital, test.ID, test.Name, test.Name, test.Capital)
		}
	}
}
