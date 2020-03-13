package main

import "fmt"

func main() {
	var i, j int

	fmt.Scanf("%d %d", &i, &j)
	k := i * j
	if k%2 != 1 {
		fmt.Print("Even")
		return
	}
	fmt.Print("Odd")
}
