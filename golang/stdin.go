package main

import "fmt"

// go run stdin.go
// 1
// 2 3
// 4 5 6
// test
func main() {
	var i1 int
	var j1, j2 int
	var k1, k2, k3 int
	var str string

	fmt.Scanf("%d", &i1)
	fmt.Scanf("%d %d", &j1, &j2)
	fmt.Scanf("%d %d %d", &k1, &k2, &k3)
	fmt.Scanf("%s", &str)

	fmt.Printf("%d\n", i1)
	fmt.Printf("%d %d\n", j1, j2)
	fmt.Printf("%d %d %d\n", k1, k2, k3)
	fmt.Printf("%s\n", str)
}
