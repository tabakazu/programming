package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/ping", PingHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PingHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!\n")
}
