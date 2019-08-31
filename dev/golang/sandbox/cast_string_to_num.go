package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v is %T\n", i, i)
	fmt.Printf("%v is %T\n\n", uint(i), uint(i))
	// => 1 is int
	// => 1 is uint

	i64, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v is %T\n", i64, i64)
	fmt.Printf("%v is %T\n", int(i64), int(i64))
	fmt.Printf("%v is %T\n\n", uint(i64), uint(i64))
	// => 1 is int64
	// => 1 is int
	// => 1 is uint

	ui64, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v is %T\n", ui64, ui64)
	fmt.Printf("%v is %T\n", uint(ui64), uint(ui64))
	fmt.Printf("%v is %T\n\n", int(ui64), int(ui64))
	// => 1 is uint64
	// => 1 is uint
	// => 1 is int
}
