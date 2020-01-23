package handler

import (
	"fmt"

	"github.com/tabakazu/awscaller/infrastructure"
	"github.com/urfave/cli"
)

type AccountHandler struct{}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) Display() cli.Command {
	return cli.Command{
		Name:  "account",
		Usage: "Shows a aws api caller account number",
		Action: func(c *cli.Context) error {
			callerID := infrastructure.GetCallerIdentity()
			fmt.Printf("- Account : %v\n", *callerID.Account)
			return nil
		},
	}
}
