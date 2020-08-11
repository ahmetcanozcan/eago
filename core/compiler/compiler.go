package compiler

import (
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

// Compiler compile scripts from string
type Compiler struct {
	vm *otto.Otto
}

// Compile file
func (c Compiler) Compile(script *Script) (*ast.Program, error) {

	prog, err := parser.ParseFile(nil, script.Filename, script.Code, 0)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, "Can not parse "+script.Filename)
	}
	return prog, nil
}

// New return new compiler
func New() *Compiler {
	return &Compiler{}
}
