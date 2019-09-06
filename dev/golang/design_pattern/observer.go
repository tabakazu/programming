package main

import "fmt"

type Subject struct {
	observers []observer
}

func (s *Subject) AddObserver(o observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notifyObservers() {
	for _, o := range s.observers {
		o.update()
	}
}

type observer interface {
	update()
}

type FirstObserver struct{}

func (o FirstObserver) update() {
	fmt.Println("FirstObserver Update!")
}

type SecondObserver struct{}

func (o SecondObserver) update() {
	fmt.Println("SecondObserver Update!")
}

type Employee struct {
	Name string
	Subject
}

func main() {
	e := Employee{Name: "A"}

	o1 := &FirstObserver{}
	o2 := &SecondObserver{}

	e.AddObserver(o1)
	e.AddObserver(o2)
	e.notifyObservers()

	return
}
