package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/common/eagofs"
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/spf13/cobra"
)

// newProjectCmd represents the newProject command
var newProjectCmd = &cobra.Command{
	Use:   "project [directory name]",
	Short: "Create new project",
	Long: `Create a new eago project in the provided directory.
	The new project contains main folders and 'hello world' files .`,
	PreRunE: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirname := args[0]
		if !filepath.IsAbs(dirname) {
			cwd, _ := os.Getwd()
			dirname = filepath.Join(cwd, dirname)
		}
		force, _ := cmd.Flags().GetBool("force")
		if err := createNewProject(dirname, force); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	newCmd.AddCommand(newProjectCmd)
}

func createNewProject(baseDir string, force bool) error {
	dirs := []string{
		filepath.Join(baseDir, "files"),
		filepath.Join(baseDir, "modules"),
		filepath.Join(baseDir, "packages"),
		filepath.Join(baseDir, "handlers"),
		filepath.Join(baseDir, "events"),
	}

	if err := runSpinnerTask("Check folder", func() error {
		if exist := eagofs.IsExist(baseDir); exist {
			if !eagofs.IsDir(baseDir) {
				return eagrors.NewError(baseDir + " already exists but not a directory")
			}
			isEmpty := eagofs.IsEmpty(baseDir)
			if !isEmpty && !force {
				return eagrors.NewError(baseDir + " already exists and is not empty. See --force.")
			} else if !isEmpty && force {
				all := append(dirs, filepath.Join(baseDir, "start.js"))
				all = append(all, filepath.Join(baseDir, "eago.json"))
				for _, path := range all {
					if eagofs.IsExist(path) {
						return eagrors.NewError(path + " already exists")
					}
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}

	if err := runSpinnerTask("Create files", func() error {
		for _, dir := range dirs {
			if err := eagofs.MkdirAll(dir, 0777); err != nil {
				return eagrors.NewErrorWithCause(err, "Failed to create dir "+dir)
			}
		}

		if err := createEagoJSON(baseDir, false); err != nil {
			return err
		}

		if err := createStartJS(baseDir); err != nil {
			return err
		}
		if err := createGitIgnore(baseDir); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
