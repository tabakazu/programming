package main

import (
	"fmt"
)

func main() {
	var n, v, cnt int
	min := 1000000000
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		cnt = f(0, v)
		if cnt < min {
			min = cnt
		}
	}
	fmt.Println(min)
}

func f(cnt int, i int) int {
	if i%2 == 1 {
		return cnt
	}
	return f(cnt+1, i/2)
}
