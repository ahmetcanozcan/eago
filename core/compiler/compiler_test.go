package compiler

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestCompilerWithoutTransformer(t *testing.T) {
	c := qt.New(t)

	compiler := &Compiler{}

	script := &Script{
		Filename: "test.js",
		Code:     `var text = "Test Message"`,
	}

	program, err := compiler.Compile(script)

	// Compiler parse code correctly
	c.Assert(err, qt.IsNil)

	vm := otto.New()

	vm.Run(program)

	textVal, _ := vm.Get("text")
	text := textVal.String()

	c.Assert(text, qt.Equals, "Test Message")

}
