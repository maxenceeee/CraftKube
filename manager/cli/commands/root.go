package commands

import "github.com/spf13/cobra"

type RootCommand struct{}

func (r *RootCommand) Execute(c *cobra.Command, args []string) error {
	return nil
}
func (r *RootCommand) Name() string {
	return "craftkube"
}
func (r *RootCommand) Description() string {
	return "CraftKube Manager CLI"
}
func (r *RootCommand) Flags() map[string]string {
	return map[string]string{}
}
