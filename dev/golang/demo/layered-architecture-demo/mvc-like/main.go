package main

import "github.com/tabakazu/layered-architecture-demo/mvc-like/config/routes"

func main() {
	r := routes.NewRoutes()
	r.Run(":8080")
}
