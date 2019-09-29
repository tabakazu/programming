package model_test

import (
	"testing"
	"time"

	"github.com/tabakazu/modeling-lesson/lesson04/model"
)

type creditcardTest struct {
	ID         int
	CustomerID int
	CardNumber string
	Expiration string
	Nominee    string
}

var (
	layout          string = "2006-01-02"
	creditcardTests        = []creditcardTest{
		{ID: 1, CustomerID: 1, CardNumber: "424242424242", Expiration: "2019-07-23", Nominee: "HANAKO SUZUKI"},
	}
)

func TestNewCreditcard(t *testing.T) {
	for i := range creditcardTests {
		test := &creditcardTests[i]
		expiration, _ := time.Parse(layout, test.Expiration)
		c := model.NewCreditcard(test.ID, test.CustomerID, test.CardNumber, expiration, test.Nominee)

		if c.ID != test.ID || c.CustomerID != test.CustomerID || c.CardNumber != test.CardNumber ||
			c.Expiration != expiration || c.Nominee != test.Nominee {

			t.Errorf("c.ID, c.CustomerID, c.CardNumber, c.Expiration, c.Nominee = %v, %v, %v, %v, %v \nwant %v, %v, %v, %v, %v",
				c.ID, c.CustomerID, c.CardNumber, c.Expiration, c.Nominee, test.ID, test.CustomerID, test.CardNumber, expiration, test.Nominee)
		}
	}
}
