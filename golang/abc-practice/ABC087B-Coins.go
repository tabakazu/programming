package main

import "fmt"

func main() {
	var a, b, c, x, cnt int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	// 500yen
	for i := 0; i <= a; i++ {
		// 100yen
		for j := 0; j <= b; j++ {
			// 50yen
			for k := 0; k <= c; k++ {
				t := 500*i + 100*j + 50*k
				if x == t {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
