package main

import "fmt"

func rec(f2, f1, n int64) int64 {
	if n == 1 {
		return f1
	} else {
		return rec(f1, f2+f1, n-1)
	}
}

func fibo(n int64) int64 {
	if n < 2 {
		return n
	} else {
		return rec(0, 1, n)
	}
}

func main() {
	fmt.Println(fibo(1000))
}
