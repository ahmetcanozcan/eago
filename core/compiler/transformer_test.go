package compiler

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTransformer(t *testing.T) {
	c := qt.New(t)
	// TODO: expand test cases
	testCases := map[string]string{
		"let a = 26": "\"use strict\";\n\nvar a = 26;",
	}

	for k, v := range testCases {
		reader := strings.NewReader(k)
		transformed, err := transformToES2015(reader)
		c.Assert(err, qt.IsNil)
		c.Assert(transformed, qt.Equals, v)
	}

}
