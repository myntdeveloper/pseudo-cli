package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [name]",
	Aliases: []string{"rm", "delete", "del"},
	Short:   "Delete a command alias by name",
	Long:    "Deletes the specified command alias from the pseudo-cli store by its name.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		return AppRunner.Remove(name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
