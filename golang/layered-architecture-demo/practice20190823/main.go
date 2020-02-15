package main

import (
	"fmt"

	"github.com/tabakazu/interface-pattern-demo/interfaces"
)

func main() {
	fmt.Println("Hello Interface Pattern")

	r := interfaces.NewRouter()
	r.Run(":8080")
}
