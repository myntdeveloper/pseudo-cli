package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show [name]",
	Aliases: []string{"sh"},
	Short:   "Show command by name",
	Long:    "Displays detailed information about a saved command alias by its name, including its description and tag if available.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		return AppRunner.Show(name)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
