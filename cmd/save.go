package cmd

import (
	"github.com/myntdeveloper/pseudo-cli/internal/models"
	"github.com/spf13/cobra"
)

var (
	saveDescription string
	saveTag         string
)

var saveCmd = &cobra.Command{
	Use:     "save [name] [command]",
	Aliases: []string{"s"},
	Short:   "Save a new pseudo command",
	Long:    "Saves a new pseudo command alias with an optional description and tag.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		command := args[1]
		pseudo := models.Pseudonym{
			Name:        name,
			Command:     command,
			Description: saveDescription,
			Tag:         saveTag,
		}
		return AppRunner.Save(pseudo)
	},
}

func init() {
	saveCmd.Flags().StringVarP(&saveDescription, "description", "d", "", "Description for the command")
	saveCmd.Flags().StringVarP(&saveTag, "tag", "t", "", "Tag for the command")
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
