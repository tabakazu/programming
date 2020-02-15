package main

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	html := HTMLFormatter{}
	plainText := PlainTextFormatter{}
	other := OtherFormatter{}

	r := Report{"Test Title", "Test Text", html}
	if r.outputReport() != "html" {
		t.Errorf("Fail")
	}

	r.Formatter = plainText
	if r.outputReport() != "plain text" {
		t.Errorf("Fail")
	}

	r.Formatter = other
	if r.outputReport() != "Abstract method called" {
		t.Errorf("Fail")
	}
}
