// https://github.com/russolsen/design_patterns_in_ruby_code/blob/master/04/ex1_strategy.rb
// https://github.com/russolsen/design_patterns_in_ruby_code/blob/master/04/ex1_strategy_demo.rb
package main

import (
	"fmt"
)

type Formatter interface {
	outputReport(title string, text string) string
}

type HTMLFormatter struct{}

func (f HTMLFormatter) outputReport(title string, text string) string {
	s := "html"
	return s
}

type PlainTextFormatter struct{}

func (f PlainTextFormatter) outputReport(title string, text string) string {
	s := "plain text"
	return s
}

type Report struct {
	Title     string
	Text      string
	Formatter Formatter
}

func (r Report) outputReport() string {
	return r.Formatter.outputReport(r.Title, r.Text)
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
	return
}
