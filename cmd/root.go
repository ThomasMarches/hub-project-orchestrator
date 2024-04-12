package cmd

import (
	"os"

	"github.com/ThomasMarches/hub-project-orchestrator.git/gen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ModuleOrchestrator",
	Short: "",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(gen.Cmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
