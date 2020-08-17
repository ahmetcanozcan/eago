package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/common/eagofs"
	"github.com/spf13/cobra"
)

// newPackageCmd represents the newPackage command
var newPackageCmd = &cobra.Command{
	Use:     "package [name]",
	Short:   "Create a new package project",
	Long:    ``,
	PreRunE: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirname := args[0]
		force, _ := cmd.Flags().GetBool("force")
		if !filepath.IsAbs(dirname) {
			cwd, _ := os.Getwd()
			dirname = filepath.Join(cwd, dirname)
		}
		if err := createNewPackage(dirname, force); err != nil {
			fmt.Println(err)
		}
	},
}

func createNewPackage(basedir string, force bool) error {

	files := []struct {
		content  string
		filename string
	}{
		{
			content:  `export default sum = (a,b) => a+b;`,
			filename: "sum.js",
		},
		{
			content: `
import sum from "./sum";

describe("Sum function tests", (test) => {
	test("4+5 is 9", (assert) => {
		assert(sum(4, 5) == 9);
	});
});
`,
			filename: "sum.spec.js",
		},
	}

	if err := runSpinnerTask("Chek folder", func() error {
		return checkBaseDir(basedir, force,
			func() []string {
				res := make([]string, len(files))
				for i := range res {
					res[i] = files[i].filename
				}
				return res
			}())
	}); err != nil {
		return err
	}

	if err := runSpinnerTask("Create files", func() error {

		if err := eagofs.MkdirAll(filepath.Join(basedir), 0777); err != nil {
			return err
		}

		for _, info := range files {
			file, err := os.Create(filepath.Join(basedir, info.filename))
			if err != nil {
				return err
			}
			_, err = file.WriteString(info.content)
			if err != nil {
				return err
			}
		}
		if err := createEagoJSON(basedir, true); err != nil {
			return err
		}
		if err := createGitIgnore(basedir); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func init() {
	newCmd.AddCommand(newPackageCmd)
}
