package cli

import (
	"github.com/spf13/cobra"
)

func initApp() {
	version := "0.1.0"
	analysisVersion := "0.0.1"

	header(version, analysisVersion)
}

func launcher(cmd *cobra.Command, args []string) {
	initApp()
}