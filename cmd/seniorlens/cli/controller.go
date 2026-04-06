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

func launcher(cmd *cobra.Command, args []string) {

}

func initFlags() {
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", ".", "set the directory of analyze")
	rootCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", true, "set recursive directory research")
}

func Controller() error {
	initFlags()

	return rootCmd.Execute()
}
