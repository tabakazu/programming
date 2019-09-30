package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson02/model"
)

var customerTests = []model.Customer{
	{ID: 1, Name: "鈴木ハナコ", Email: "hanako@example.jp"},
	{ID: 2, Name: "佐藤イチロー", Email: "ichiro@example.jp"},
}

func TestNewCustomer(t *testing.T) {
	for i := range customerTests {
		test := &customerTests[i]
		c := model.NewCustomer(test.ID, test.Name, test.Email)
		if c.ID != test.ID {
			t.Errorf("c.ID = %v want %v", c.ID, test.ID)
		}
		if c.Name != test.Name {
			t.Errorf("c.Name = %v want %v", c.Name, test.Name)
		}
		if c.Email != test.Email {
			t.Errorf("c.Email = %v want %v", c.Email, test.Email)
		}
	}
}
