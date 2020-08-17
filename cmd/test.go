package cmd

import (
	"os"

	"github.com/ahmetcanozcan/eago/core"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs all test files",
	Long:  `Eago packages can be tested as every software`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		core.Test(cwd)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
