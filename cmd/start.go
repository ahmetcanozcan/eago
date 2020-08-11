package cmd

import (
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/core"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start eago",
	Long: `
		Starts your eago application
`,
	Run: func(cmd *cobra.Command, args []string) {
		dirFlag := cmd.Flag("dir").Value.String()
		path, _ := os.Getwd()
		if dirFlag != "" {
			path = filepath.Join(path, dirFlag)
		}
		core.Start(core.StartOptions{AppDir: path})
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().String("dir", "", "source directory of application")
}
