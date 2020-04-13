package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("index.html.tpl"))
		m := map[string]string{
			"Hello": "Hello World!",
		}
		tpl.Execute(w, m)
	})

	http.ListenAndServe(":8080", nil)
}
