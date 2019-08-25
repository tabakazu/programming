package main

import "github.com/tabakazu/goapp-practice/interfaces/api"

func main() {
	r := api.NewRouter()
	r.Run(":8080")
}
