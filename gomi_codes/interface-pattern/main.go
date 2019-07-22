package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/interface-pattern/interfaces/http"
)

func main() {
	fmt.Println("Hello Interface Pattern")

	r := http.NewRouter()
	r.Run(":8080")
}
