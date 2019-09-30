package model_test

import (
	"testing"
	"time"

	"github.com/tabakazu/modeling-lesson/lesson05/model"
)

type regularCustomerTest struct {
	ID       int
	Name     string
	Email    string
	Sex      int
	Birthday string
}

var (
	regularCustomerTests = []regularCustomerTest{
		{ID: 1, Name: "鈴木ハナコ", Email: "hanako@example.jp", Sex: 1, Birthday: "1999-09-09"},
		{ID: 2, Name: "佐藤イチロー", Email: "ichiro@example.jp", Sex: 2, Birthday: "2000-02-22"},
	}
)

func TestNewRegularCustomer(t *testing.T) {
	for i := range regularCustomerTests {
		test := &regularCustomerTests[i]
		birthday, _ := time.Parse("2006-01-02", test.Birthday)
		c := model.NewRegularCustomer(test.ID, test.Name, test.Email, test.Sex, birthday)

		if c.ID != test.ID || c.Name != test.Name || c.Email != test.Email ||
			c.Sex != test.Sex || c.Birthday != birthday {

			t.Errorf("c.ID, c.Name, c.Email, c.Sex, c.Birthday = %v, %v, %v, %v, %v want %v, %v, %v, %v, %v",
				c.ID, c.Name, c.Email, c.Sex, c.Birthday, test.ID, test.Name, test.Name, test.Sex, birthday)
		}
	}
}
