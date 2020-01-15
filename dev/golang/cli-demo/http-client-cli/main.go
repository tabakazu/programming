package main

import (
	"fmt"

	"github.com/tabakazu/cli-demo/http-client-cli/interface/cli"
)

func main() {
	url := "https://google.com"

	cmd := cli.NewCli()
	site, err := cmd.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(url, site.StatusCode)
}
