package model_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson03/model"
)

type addressTest struct {
	ID            int
	CustomerID    int
	Name          string
	StreetAddress string
}

var addressTests = []addressTest{
	{1, 1, "自宅", "中央区銀座1-1"},
	{2, 1, "職場", "港区六本木1-1"},
}

func TestNewAddress(t *testing.T) {
	for i := range addressTests {
		test := &addressTests[i]
		a := model.NewAddress(test.ID, test.CustomerID, test.Name, test.StreetAddress)
		if a.ID != test.ID || a.CustomerID != test.CustomerID ||
			a.Name != test.Name || a.StreetAddress != test.StreetAddress {

			t.Errorf("a.ID, a.CustomerID, a.Name, a.StreetAddress = %v, %v, %v, %v \nwant %v, %v, %v, %v",
				a.ID, a.CustomerID, a.Name, a.StreetAddress, test.ID, a.CustomerID, test.Name, test.StreetAddress)
		}
	}
}
