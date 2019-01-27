package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mycli"

	app.Commands = []cli.Command{
		{
			Name: "hello",
			Aliases: []string{"he"},
			Usage: "say hello to NAME",
			Action: func(c *cli.Context) error {
				fmt.Println("Hello", c.Args().First())
				return nil
			},
		},
	}

  app.Run(os.Args)
}
