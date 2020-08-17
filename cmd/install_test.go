package cmd

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestVerifyPackageName(t *testing.T) {
	c := qt.New(t)

	testCases := []struct {
		name     string
		expected bool
	}{
		{"github.com/organization/repo", true},
		{"http://github.com/organization/repo", false},
	}

	for _, _case := range testCases {
		c.Assert(verifyPackageName(_case.name), qt.Equals, _case.expected)
	}
}
