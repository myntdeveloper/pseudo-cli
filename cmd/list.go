package cmd

import (
	"github.com/spf13/cobra"
)

var (
	listTag string
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List saved pseudo commands",
	Long:    "Lists all saved pseudo command aliases. Optionally filter the list by tag using --tag or -t.",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return AppRunner.List(listTag)
	},
}

func init() {
	listCmd.Flags().StringVarP(&listTag, "tag", "t", "", "Filter pseudonyms by tag")
}

func init() {
	rootCmd.AddCommand(listCmd)
}
