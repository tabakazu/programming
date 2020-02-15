package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Printf("%d\n", strings.Count(str, "1"))
}
