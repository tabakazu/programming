package main

import (
	"os"

	"github.com/tabakazu/awscaller/interfaces"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "awscaller"
	app.Usage = "Shows aws api caller infomations"
	app.Version = "0.0.4"
	app.Commands = interfaces.NewCommands()
	app.Run(os.Args)
}
