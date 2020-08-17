package lib

import "path/filepath"

var (
	// ModuleDirPath is modules directory of project
	ModuleDirPath string
	// HandlerDirPath is handlers directory of project
	HandlerDirPath string
	// NitDirPath is nits directory of project
	NitDirPath string
	// FileDirPath :
	FileDirPath string
	// PackageDirPath is the directory that contains 3rd party modules
	PackageDirPath string
	// BasePath is the main directory of application
	BasePath string
)

// UpdateEnginePathVars update path variables depends on given base dir
func UpdateEnginePathVars(baseDir string) {
	BasePath = baseDir
	ModuleDirPath = filepath.Join(baseDir, "modules")
	HandlerDirPath = filepath.Join(baseDir, "handlers")
	NitDirPath = filepath.Join(baseDir, "nits")
	FileDirPath = filepath.Join(baseDir, "files")
	PackageDirPath = filepath.Join(baseDir, "packages")
}
