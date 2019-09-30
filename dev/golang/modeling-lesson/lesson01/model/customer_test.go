package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson01/model"
)

var customerTests = []model.Customer{
	{ID: 1},
	{ID: 100},
}

func TestNewCustomer(t *testing.T) {
	for i := range customerTests {
		test := &customerTests[i]
		c := model.NewCustomer(test.ID)
		if c.ID != test.ID {
			t.Errorf("c.ID = %v want %v", c.ID, test.ID)
		}
	}
}
