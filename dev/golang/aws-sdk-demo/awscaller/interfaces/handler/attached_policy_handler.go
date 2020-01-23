package handler

import (
	"fmt"
	"regexp"

	"github.com/tabakazu/awscaller/infrastructure"
	"github.com/urfave/cli"
)

type AttachedPolicyHandler struct{}

func NewAttachedPolicyHandler() *AttachedPolicyHandler {
	return &AttachedPolicyHandler{}
}

func (h *AttachedPolicyHandler) DisplayList() cli.Command {
	return cli.Command{
		Name:  "policies",
		Usage: "Shows a aws api caller attached policies",
		Action: func(c *cli.Context) error {
			callerID := infrastructure.GetCallerIdentity()
			userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
			listPolicies := infrastructure.GetListAttachedUserPolicies(userName)
			fmt.Println("- AttachedPolicies :")
			for _, policy := range listPolicies.AttachedPolicies {
				fmt.Printf("\t- %v\n", *policy.PolicyName)
			}
			return nil
		},
	}
}
