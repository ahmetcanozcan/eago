package core

import (
	"testing"

	"github.com/ahmetcanozcan/eago/core/compiler"
	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestEvaluateScriptToBundles(t *testing.T) {
	c := qt.New(t)
	const testKey = "lib/example.js"
	scripts := map[string]*compiler.Script{
		testKey: {
			Filename: "lib/example.js",
			Code: `
				var testString = "testString"
			`,
		},
	}

	bundles, err := evaluateSciptsToBundles(scripts)

	c.Assert(err, qt.IsNil)

	vm := otto.New()
	vm.Run(bundles[testKey].Program)
	r, _ := vm.Get("testString")
	str := r.String()
	c.Assert(str, qt.Equals, "testString")
}
