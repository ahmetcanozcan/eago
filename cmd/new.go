package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [command]",
	Short: "Create projects and packages ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	newCmd.PersistentFlags().BoolP("force", "f", false, "Force to create")
	rootCmd.AddCommand(newCmd)

}
