package core

import (
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/common/eagofs"
	"github.com/ahmetcanozcan/eago/core/compiler"
	"github.com/ahmetcanozcan/eago/core/lib"
)

func parseScriptsFromDir(dirname string) (map[string]*compiler.Script, error) {
	result := make(map[string]*compiler.Script)
	err := eagofs.NewDirReader(dirname).Read(func(filename string) error {
		// If this file is not a JS file, ignore it.
		if !strings.HasSuffix(filename, ".js") {
			return nil
		}
		script, err := compiler.ReadScript(filepath.Join(dirname, filename))
		if err != nil {
			return err
		}
		result[filename] = script
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, err
}

func evaluateSciptsToBundles(scripts map[string]*compiler.Script) (map[string]*lib.Bundle, error) {
	bundles := make(map[string]*lib.Bundle)
	comp := compiler.New()
	for key, script := range scripts {
		prog, err := comp.Compile(script)
		if err != nil {
			return nil, err
		}
		bundles[key] = lib.NewBundle(
			prog,
			lib.NewProgramPath(key),
		)
	}
	return bundles, nil
}
