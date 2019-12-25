package main_test

import (
	"testing"
	"time"

	"github.com/tabakazu/modeling-lesson/lesson12/model"
)

type airplaneFlightTest struct {
	ID                int
	Name              string
	DepartureLocation string
	DepartedAt        time.Time
	ArrivalLocation   string
	ArrivedAt         time.Time
}

var t1 time.Time = time.Now()
var t2 time.Time = t1.Add(1 * time.Hour)

var airplaneFlightTests = []airplaneFlightTest{
	{Name: "東京便", DepartureLocation: "大阪", DepartedAt: t1, ArrivalLocation: "東京", ArrivedAt: t2},
}

func TestMain(t *testing.T) {
	for i := range airplaneFlightTests {
		test := &airplaneFlightTests[i]
		af := model.NewAirplaneFlight(test.Name, test.DepartureLocation, test.DepartedAt, test.ArrivalLocation, test.ArrivedAt)

		if af.Name != test.Name || af.DepartureLocation != test.DepartureLocation || af.DepartedAt != test.DepartedAt ||
			af.ArrivalLocation != test.ArrivalLocation || af.ArrivedAt != test.ArrivedAt {
			t.Errorf("af.Name, af.DepartureLocation, af.DepartedAt, af.ArrivalLocation, af.ArrivedAt = %v, %v, %v, %v, %v want %v, %v, %v, %v, %v ",
				af.Name, af.DepartureLocation, af.DepartedAt, af.ArrivalLocation, af.ArrivedAt, test.Name, test.DepartureLocation, test.DepartedAt, test.ArrivalLocation, test.ArrivedAt)
		}
	}
}
