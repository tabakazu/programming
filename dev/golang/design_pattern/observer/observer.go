package main

import "fmt"

type Subject struct {
	observers []observer
}

func (s *Subject) AddObserver(o observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notifyObservers() string {
	for _, o := range s.observers {
		o.update()
	}
	str := "success"
	return str
}

type observer interface {
	update() string
}

type FirstObserver struct{}

func (o FirstObserver) update() string {
	s := "FirstObserver Update!"
	fmt.Println(s)
	return s
}

type SecondObserver struct{}

func (o SecondObserver) update() string {
	s := "SecondObserver Update!"
	fmt.Println(s)
	return s
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
