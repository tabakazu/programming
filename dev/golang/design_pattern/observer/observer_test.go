package main

import (
	"testing"
)

func TestObserver(t *testing.T) {
	e := Employee{Name: "A"}

	o1 := &FirstObserver{}
	o2 := &SecondObserver{}

	e.AddObserver(o1)
	e.AddObserver(o2)
	if e.notifyObservers() != "success" {
		t.Errorf("Fail")
	}
}
