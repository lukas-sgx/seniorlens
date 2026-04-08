package cli

import (
	"github.com/spf13/cobra"
)

var (
	directory string
	recursive bool
)

var rootCmd = &cobra.Command{
	Use:   "seniorlens",
	Short: "SeniorLens CLI tool",
	Run:   launcher,
}

func initFlags() {
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "set the directory of analyze")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", true, "set recursive directory research")
}

func Controller() error {
	initFlags()

	return rootCmd.Execute()
}
