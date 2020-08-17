package lib

import (
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
)

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
