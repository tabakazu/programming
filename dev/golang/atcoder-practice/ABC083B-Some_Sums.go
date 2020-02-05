package main

import "fmt"

func main() {
	var n, a, b, total int
	fmt.Scan(&n, &a, &b)
	for i := 1; i <= n; i++ {
		v := f(i)
		if a <= v && v <= b {
			total = total + i
		}
	}
	fmt.Println(total)
}

func f(n int) int {
	var total int
	for n > 0 {
		total = total + n%10
		n = n / 10
	}
	return total
}
