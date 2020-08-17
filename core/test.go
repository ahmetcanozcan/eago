package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/common/loggers"
	"github.com/ahmetcanozcan/eago/config"
	"github.com/ahmetcanozcan/eago/core/engine"
	"github.com/ahmetcanozcan/eago/core/engine/js/bootstrap"
	"github.com/ahmetcanozcan/eago/core/lib"
)

// Test run all tests on package directory.
func Test(basedir string) {

	config.Parse(basedir)
	loggers.InitializeLoggers()
	lib.UpdateEnginePathVars(basedir)

	bundles, err := parseSpecFiles(basedir)
	if err != nil {
		loggers.Default().Error(err)
	}
	for _, bundle := range bundles {
		vm := engine.GetBaseRuntime()
		dirname, _ := filepath.Split(string(bundle.Filename))
		dirname = strings.ReplaceAll("@/"+filepath.Join(config.EagoJSON.Name, dirname)+"/", "\\", "/")
		fmt.Println("dirname", dirname)
		vm.Set("__dirname", dirname)
		vm.Run(bootstrap.TestBootstrapProgram)
		_, err := vm.Run(bundle.Program)
		if err != nil {
			loggers.Default().Error(err)
		}
	}
}

func parseSpecFiles(basedir string) ([]*lib.Bundle, error) {
	scripts, err := parseScriptsFromDir(basedir)
	if err != nil {
		return nil, err
	}
	bundles, err := evaluateSciptsToBundles(scripts)
	if err != nil {
		return nil, err
	}
	res := make([]*lib.Bundle, 0)
	for _, bundle := range bundles {
		if strings.HasSuffix(string(bundle.Filename), ".spec.js") {
			res = append(res, bundle)
		}
	}

	return res, nil
}
