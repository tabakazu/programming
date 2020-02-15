package handler

import (
	"fmt"

	"github.com/tabakazu/awscaller/usecase"
	"github.com/urfave/cli"
)

type CallerHandler struct{}

func NewCallerHandler() *CallerHandler {
	return &CallerHandler{}
}

func (h *CallerHandler) DisplayAccountId() cli.Command {
	u := usecase.NewCallerUsecase()
	str, _ := u.GetAccountId()

	return cli.Command{
		Name:  "account_id",
		Usage: "Shows a aws api caller account number",
		Action: func(c *cli.Context) error {
			fmt.Print(str)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayUserId() cli.Command {
	u := usecase.NewCallerUsecase()
	str, _ := u.GetUserId()

	return cli.Command{
		Name:  "user_id",
		Usage: "Shows a aws api caller user id",
		Action: func(c *cli.Context) error {
			fmt.Print(str)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayUserName() cli.Command {
	u := usecase.NewCallerUsecase()
	str, _ := u.GetUserName()

	return cli.Command{
		Name:  "username",
		Usage: "Shows a aws api caller user name",
		Action: func(c *cli.Context) error {
			fmt.Print(str)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayAttachedUserPolicies() cli.Command {
	u := usecase.NewCallerUsecase()
	userName, _ := u.GetUserName()
	strArray, _ := u.GetAttachedUserPolicies(userName)

	return cli.Command{
		Name:  "policies",
		Usage: "Shows a aws api caller attached policies",
		Action: func(c *cli.Context) error {
			for _, v := range strArray {
				fmt.Printf("- %v\n", v)
			}
			return nil
		},
	}
}
