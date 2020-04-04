package main

import (
	"fmt"

	"github.com/tabakazu/gortfolio/infrastructure"
)

func main() {
	infrastructure.Web.Run(":8080")
	fmt.Println("Go Application starting...")
}
