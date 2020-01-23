package interfaces

import (
	"github.com/tabakazu/awscaller/interfaces/handler"
	"github.com/urfave/cli"
)

func NewCommands() []cli.Command {
	account := handler.NewAccountHandler()
	attachedPolicy := handler.NewAttachedPolicyHandler()
	user := handler.NewUserHandler()
	commands := []cli.Command{
		account.Display(),
		attachedPolicy.DisplayList(),
		user.DisplayId(),
		user.DisplayName(),
	}
	return commands
}
