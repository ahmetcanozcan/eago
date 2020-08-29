package lib

import "path/filepath"

var (
	// ModuleDirPath is modules directory of project
	ModuleDirPath string
	// HandlerDirPath is handlers directory of project
	HandlerDirPath string
	// EventDirPath is events directory of project
	EventDirPath string
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
	EventDirPath = filepath.Join(baseDir, "events")
	FileDirPath = filepath.Join(baseDir, "files")
	PackageDirPath = filepath.Join(baseDir, "packages")
}
