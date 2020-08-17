package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/config"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install [package name]",
	Aliases: []string{"i"},
	Short:   "Install package",
	Long: `
	Instal packages using git `,
	PreRunE: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		save, _ := cmd.Flags().GetBool("save")
		saveDev, _ := cmd.Flags().GetBool("save-dev")
		if err := installPackage(args[0], save, saveDev); err != nil {
			fmt.Println(err)
		}
	},
}

func installPackage(name string, save, saveDev bool) error {
	cwd, _ := os.Getwd()

	if err := runSpinnerTask("Verify package name", func() error {
		// Read eago.json
		config.Parse(cwd)
		return verifyPackageName(name)
	}); err != nil {
		return err
	}

	if err := runSpinnerTask("Install package", func() error {
		return exec.Command("git", "clone", "https://"+name, filepath.Join("packages", name)).Run()
	}); err != nil {
		return err
	}

	if save || saveDev {
		if err := runSpinnerTask("Update eago.json", func() error {

			// set dependincy and save it
			if save {
				config.EagoJSON.Dependincies[name] = "latest"
			}
			if saveDev {
				config.EagoJSON.DevDependincies[name] = "latest"
			}
			return config.EagoJSON.Save()
		}); err != nil {
			return err
		}
	}

	return nil
}

func verifyPackageName(name string) error {
	// If package name contains not allowed parts, return err
	for _, part := range []string{"//"} { // Not allowed parts
		if strings.Contains(name, part) {
			return errors.New("Invalid Part :" + part)
		}
	}
	return nil
}

func init() {
	installCmd.PersistentFlags().BoolP("save", "s", false, "add package to  dependincies")
	installCmd.PersistentFlags().BoolP("save-dev", "D", false, "add package to devDependincies")
	rootCmd.AddCommand(installCmd)
}
