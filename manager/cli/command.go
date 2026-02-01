package cli

import (
	"github.com/spf13/cobra"
	"xamence.eu/craftkube/manager/cli/commands"
	"xamence.eu/craftkube/manager/repository"
)

var Commands []CLICommand
var cobraRootCmd *cobra.Command

type CLICommand interface {
	Execute(c *cobra.Command, args []string) error
	Name() string
	Description() string
	Flags() map[string]string
}

func registerCommand(cmd CLICommand) {
	Commands = append(Commands, cmd)

	cobraCmd := &cobra.Command{
		Use:   cmd.Name(),
		Short: cmd.Description(),
		RunE: func(c *cobra.Command, args []string) error {
			return cmd.Execute(c, args)
		},
	}

	for flag, desc := range cmd.Flags() {
		cobraCmd.Flags().String(flag, "", desc)
	}

	if cobraRootCmd == nil && cmd.Name() == "craftkube" {
		cobraRootCmd = cobraCmd
	} else {
		cobraRootCmd.AddCommand(cobraCmd)
	}
}

func InitCommands(repo *repository.ServiceRepository) {
	registerCommand(&commands.RootCommand{})
	registerCommand(&commands.ServiceCommand{RepoService: repo})
}

func GetRootCommand() *cobra.Command {
	return cobraRootCmd
}
