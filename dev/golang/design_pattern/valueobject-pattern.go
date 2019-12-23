package main

import (
	"fmt"
)

type Money struct {
	Amount int
}

func (m Money) Equals(obj *Dollar) bool {
	return m.Amount == obj.Amount
}

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	d := &Dollar{}
	d.Amount = amount
	return d
}

func (d Dollar) Times(multiplier int) *Dollar {
	amount := d.Amount * multiplier
	obj := &Dollar{}
	obj.Amount = amount
	return obj
}

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	f := &Franc{}
	f.Amount = amount
	return f
}

func (f Franc) Times(multiplier int) *Franc {
	amount := f.Amount * multiplier
	obj := &Franc{}
	obj.Amount = amount
	return obj
}

func main() {
	d := NewDollar(5)
	fmt.Println(d.Times(2).Amount)
	fmt.Println(d.Equals(NewDollar(10)))

	f := NewFranc(5)
	fmt.Println(f.Times(2).Amount)
	fmt.Println(f.Equals(d))
}
