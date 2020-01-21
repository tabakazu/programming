package main

import "github.com/tabakazu/golang-api-srv/interfaces"

func main() {
	r := interfaces.NewRouter()
	r.Run(":3000")
}
