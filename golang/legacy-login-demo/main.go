package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatal(err)
		}

		token := genCSRFToken()
		if err := t.Execute(w, token); err != nil {
			log.Fatal(err)
		}

	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// CSRF Token 検証
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	}
}

func genCSRFToken() string {
	h := md5.New()
	crutime := time.Now().Unix()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	io.WriteString(h, "xxxxxxxxxxxxx")
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/login", LoginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
