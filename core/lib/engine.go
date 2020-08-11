package lib

import (
	"path/filepath"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
)

var (
	// ModuleDirPath is modules directory of project
	ModuleDirPath string
	// HandlerDirPath is handlers directory of project
	HandlerDirPath string
	// NitDirPath is nits directory of project
	NitDirPath string
	// FileDirPath :
	FileDirPath string
)

// UpdateEnginePathVars update path variables depends on given base dir
func UpdateEnginePathVars(baseDir string) {
	ModuleDirPath = filepath.Join(baseDir, "modules")
	HandlerDirPath = filepath.Join(baseDir, "handlers")
	NitDirPath = filepath.Join(baseDir, "nits")
	FileDirPath = filepath.Join(baseDir, "files")
}

// ProgramPath is path of a progrm
type ProgramPath string

// NewProgramPath returns a new ProgramPath
func NewProgramPath(path string) ProgramPath {
	return ProgramPath(path)
}

// Bundle is parsed javascript files ready to execution
type Bundle struct {
	Program  *ast.Program
	Filename ProgramPath
}

// NewBundle creates new bundle
func NewBundle(program *ast.Program, filename ProgramPath) *Bundle {
	return &Bundle{program, filename}
}

// JSFunc presents javascript function
type JSFunc func(call otto.FunctionCall) otto.Value
