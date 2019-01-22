package main

import (
	"flag"
	"fmt"
	"time"
	"net/http"
)

func main() {
	flag.Parse()

	res, err := http.Get(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v %v\n", res.Proto, res.Status)
	fmt.Printf("Date: %v\n", time.Now().Format("2006-01-02 15:04:05"))
}