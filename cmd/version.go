package cmd

import (
	"fmt"

	"github.com/ahmetcanozcan/eago/common/eago"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of Eeago software",
	Long:  `Full detailed version of eago software`,
	Run: func(cmd *cobra.Command, args []string) {
		version := eago.Info.BuildVersionString()
		fmt.Println(version)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
