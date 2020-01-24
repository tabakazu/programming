package handler

import (
	"fmt"

	"github.com/tabakazu/awscaller/infrastructure/awsapi"
	"github.com/urfave/cli"
)

type CallerHandler struct{}

func NewCallerHandler() *CallerHandler {
	return &CallerHandler{}
}

func (h *CallerHandler) DisplayAccountId() cli.Command {
	r := awsapi.NewCallerResource()
	clr, _ := r.GetIdentity()

	return cli.Command{
		Name:  "accountId",
		Usage: "Shows a aws api caller account number",
		Action: func(c *cli.Context) error {
			fmt.Print(clr.AccountId)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayUserId() cli.Command {
	r := awsapi.NewCallerResource()
	clr, _ := r.GetIdentity()

	return cli.Command{
		Name:  "userId",
		Usage: "Shows a aws api caller user id",
		Action: func(c *cli.Context) error {
			fmt.Print(clr.UserId)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayUserName() cli.Command {
	r := awsapi.NewCallerResource()
	clr, _ := r.GetIdentity()

	return cli.Command{
		Name:  "username",
		Usage: "Shows a aws api caller user name",
		Action: func(c *cli.Context) error {
			fmt.Print(clr.UserName)
			return nil
		},
	}
}

func (h *CallerHandler) DisplayAttachedUserPolicies() cli.Command {
	r := awsapi.NewCallerResource()
	i, _ := r.GetIdentity()
	clr, _ := r.GetAttachedIamPoliciesByUserName(i.UserName)

	return cli.Command{
		Name:  "policies",
		Usage: "Shows a aws api caller attached policies",
		Action: func(c *cli.Context) error {
			for _, v := range clr.AttachedUserPolicies {
				fmt.Printf("- %v\n", v)
			}
			return nil
		},
	}
}
