package main

import (
	"github.com/tabakazu/simple-auth-server-demo/adapter/webapi"
)

func main() {
	r := webapi.SetRoutes()
	r.Run(":8080")
}
