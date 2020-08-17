package cmd

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestVerifyPackageName(t *testing.T) {
	c := qt.New(t)

	testCases := []struct {
		name  string
		isNil bool
	}{
		{"github.com/organization/repo", true},
		{"http://github.com/organization/repo", false},
	}

	for _, _case := range testCases {
		if _case.isNil {
			c.Assert(verifyPackageName(_case.name), qt.IsNil)
		} else {
			c.Assert(verifyPackageName(_case.name), qt.Not(qt.IsNil))
		}
	}

}
