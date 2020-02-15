// main_test.go
package main_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson11/model"
)

type customerCompanyTest struct {
	ID          int
	Name        string
	Location    string
	Capital     int
	CreditLimit int
}

var customerCompanyTests = []customerCompanyTest{
	{Name: "12345", Location: "ぼっちゃん", Capital: 1, CreditLimit: 1},
}

func TestMain(t *testing.T) {
	for i := range customerCompanyTests {
		test := &customerCompanyTests[i]
		cc := model.NewCustomerCompany(test.Name, test.Location, test.Capital, test.CreditLimit)

		if cc.Name != test.Name || cc.Location != test.Location ||
			cc.Capital != test.Capital || cc.CreditLimit != test.CreditLimit {
			t.Errorf("cc.Name, cc.Location, cc.Capital, cc.CreditLimit = %v, %v, %v, %v want %v, %v, %v, %v ",
				cc.Name, cc.Location, cc.Capital, cc.CreditLimit, test.Name, test.Location, test.Capital, test.CreditLimit)
		}
	}
}
