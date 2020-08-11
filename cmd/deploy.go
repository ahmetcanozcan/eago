package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your Eago app",
	Long: `
	Deploy your Eago app on given deploy using ... 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
