package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [name]",
	Aliases: []string{"r"},
	Short:   "Run a saved pseudo command",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		p, err := AppRunner.Store.GetByName(name)
		if err != nil {
			return err
		}

		parts := strings.Fields(p.Command)
		c := exec.Command(parts[0], parts[1:]...)

		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin

		return c.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
