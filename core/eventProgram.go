package core

import (
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/core/compiler"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto/ast"
)

func getEventPrograms(basedir string) (map[string]*ast.Program, error) {
	res := make(map[string]*ast.Program)

	scripts, err := parseScriptsFromDir(lib.EventDirPath)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err,
			"can not parse events")
	}
	c := compiler.New()
	for filename, script := range scripts {
		name := filename[1 : len(filename)-len(".js")]
		program, err := c.Compile(script)
		if err != nil {
			return nil, err
		}
		res[name] = program
	}
	return res, nil
}
