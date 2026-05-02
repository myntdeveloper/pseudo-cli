package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/myntdeveloper/pseudo-cli/internal/runner"
	"github.com/myntdeveloper/pseudo-cli/internal/store"
	"github.com/spf13/cobra"
)

var AppRunner *runner.Runner

func getDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".pseudo-cli", "app.db"), nil
}

var rootCmd = &cobra.Command{
	Use:   "psd",
	Short: "Pseudo CLI - easily create and manage pseudo alias commands",
	Long: `Pseudo CLI (psd) lets you save, list, run, and manage custom shell commands using short aliases.
You can tag commands, provide descriptions, and reuse parameterized templates for easy, efficient workflows.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dbPath, err := getDBPath()
		if err != nil {
			return fmt.Errorf("failed to determine data directory: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
			return fmt.Errorf("failed to create data directory: %w", err)
		}

		s, err := store.New(dbPath)
		if err != nil {
			return fmt.Errorf("failed to initialize data store: %w", err)
		}
		AppRunner = &runner.Runner{Store: s}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
