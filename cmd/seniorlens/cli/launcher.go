package cli

import (
	"github.com/lukas-sgx/seniorlens/internal/analyzer"
	"github.com/spf13/cobra"
)



func launcher(cmd *cobra.Command, args []string) {
	initApp()
	
	analyzer.CodeReport(directory, recursive)
}