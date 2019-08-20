package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Username string `short:"u" long:"username" description:"username"`
	Path     string `short:"p" long:"path" description:"path to use"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		return
	}
	fmt.Println(opts.Username)
	fmt.Println(opts.Path)
}
