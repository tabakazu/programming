package main

import (
	"fmt"

	"github.com/tabakazu/gortfolio/infrastructure/web"
)

func main() {
	web.Router.Run(":8080")
	fmt.Println("Go Application starting...")
}
