// https://github.com/russolsen/design_patterns_in_ruby_code/blob/master/04/ex1_strategy.rb
// https://github.com/russolsen/design_patterns_in_ruby_code/blob/master/04/ex1_strategy_demo.rb
package main

import (
	"errors"
	"fmt"
)

type IFormatter interface {
	outputReport(title string, text string) (string, error)
}

type Formatter struct{}

func (f Formatter) outputReport(title string, text string) (string, error) {
	err := errors.New("Abstract method called")
	return "", err
}

type HTMLFormatter struct {
	Formatter
}

func (f HTMLFormatter) outputReport(title string, text string) (string, error) {
	s := "html"
	return s, nil
}

type PlainTextFormatter struct {
	Formatter
}

func (f PlainTextFormatter) outputReport(title string, text string) (string, error) {
	s := "plain text"
	return s, nil
}

type OtherFormatter struct {
	Formatter
}

type Report struct {
	Title     string
	Text      string
	Formatter IFormatter
}

func (r Report) outputReport() string {
	s, err := r.Formatter.outputReport(r.Title, r.Text)
	if err != nil {
		return err.Error()
	}
	return s
}

func main() {
	r := Report{
		Title:     "TEST TITLE",
		Text:      "TEST BODY",
		Formatter: HTMLFormatter{},
	}
	fmt.Println(r.outputReport())

	r.Formatter = PlainTextFormatter{}
	fmt.Println(r.outputReport())

	r.Formatter = OtherFormatter{}
	fmt.Println(r.outputReport())

	return
}
