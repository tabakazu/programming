// Stdin
package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Println("数字を入力してください")
	fmt.Scan(&i)
	fmt.Printf("%v is %T\n\n", i, i)

	var s string
	fmt.Println("文字を入力してください")
	fmt.Scan(&s)
	fmt.Printf("%v is %T\n\n", s, s)
}
