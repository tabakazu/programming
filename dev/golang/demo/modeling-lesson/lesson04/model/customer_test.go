package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson04/model"
)

type customerTest struct {
	ID    int
	Name  string
	Email string
}

var customerTests = []customerTest{
	{ID: 1, Name: "鈴木ハナコ", Email: "hanako@example.jp"},
	{ID: 2, Name: "佐藤イチロー", Email: "ichiro@example.jp"},
}

func TestNewCustomer(t *testing.T) {
	for i := range customerTests {
		test := &customerTests[i]
		c := model.NewCustomer(test.ID, test.Name, test.Email)

		if c.ID != test.ID || c.Name != test.Name || c.Email != test.Email {
			t.Errorf("c.ID, c.Name, c.Email = %v, %v, %v want %v, %v, %v", c.ID, c.Name, c.Email, test.ID, test.Name, test.Name)
		}
	}
}
