package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/common/eagofs"
	"github.com/ahmetcanozcan/eago/common/eagrors"
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

func createEagoJSON(dir string, isPackage bool) error {
	_, name := filepath.Split(dir)

	_json := []byte(fmt.Sprintf(`{
	"name": "%s",
	"version": "0.1.0",
	"package": %t,
	"port": 3000,
	"notFound": "404.html",
	"dependincies": {},
	"devDependincies": {}
}`, name, isPackage))

	file, err := os.Create(filepath.Join(dir, "eago.json"))
	if err != nil {
		return err
	}
	_, err = file.Write(_json)
	if err != nil {
		return err
	}
	return nil
}

func createStartJS(basedir string) error {
	return eagofs.CreateFile(filepath.Join(basedir, "start.js"),
		`const msg = "Eago Application started.";
	console.log(msg);`)
}

func createGitIgnore(basedir string) error {

	return eagofs.CreateFile(filepath.Join(basedir, ".gitignore"),
		`/packages/
*.log
	`)
}

func checkBaseDir(dirname string, force bool, parts []string) error {
	if exist := eagofs.IsExist(dirname); exist {
		if !eagofs.IsDir(dirname) {
			return eagrors.NewError(dirname + " already exists but not a directory")
		}
		isEmpty := eagofs.IsEmpty(dirname)
		if !isEmpty && !force {
			return eagrors.NewError(dirname + " already exists and is not empty. See --force.")
		} else if !isEmpty && force {
			for _, path := range parts {
				if eagofs.IsExist(path) {
					return eagrors.NewError(path + " already exists")
				}
			}
		}
	}
	return nil
}
