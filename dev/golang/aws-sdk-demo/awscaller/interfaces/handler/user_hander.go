package handler

import (
	"fmt"
	"regexp"

	"github.com/tabakazu/awscaller/infrastructure"
	"github.com/urfave/cli"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) DisplayId() cli.Command {
	return cli.Command{
		Name:  "userid",
		Usage: "Shows a aws api caller user id",
		Action: func(c *cli.Context) error {
			callerID := infrastructure.GetCallerIdentity()
			fmt.Printf("- UserId : %v\n", *callerID.UserId)
			return nil
		},
	}
}

func (h *UserHandler) DisplayName() cli.Command {
	return cli.Command{
		Name:  "username",
		Usage: "Shows a aws api caller user name",
		Action: func(c *cli.Context) error {
			callerID := infrastructure.GetCallerIdentity()
			userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
			fmt.Printf("- UserName : %v\n", userName)
			return nil
		},
	}
}
