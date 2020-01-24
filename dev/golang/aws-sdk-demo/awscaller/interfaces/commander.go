package interfaces

import (
	"github.com/tabakazu/awscaller/interfaces/handler"
	"github.com/urfave/cli"
)

func NewCommands() []cli.Command {
	h := handler.NewCallerHandler()
	commands := []cli.Command{
		h.DisplayAccountId(),
		h.DisplayUserId(),
		h.DisplayUserName(),
		h.DisplayAttachedUserPolicies(),
	}
	return commands
}
