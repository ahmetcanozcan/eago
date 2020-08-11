package eago

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestVersion(t *testing.T) {
	c := qt.New(t)
	v1 := newVersion(1, 1, 5, devVersionSuffix).String()
	v2 := newVersion(0, 2, 3, emptyVersionSuffix).String()

	c.Assert(v1, qt.Equals, "1.1.5-DEV")
	c.Assert(v2, qt.Equals, "0.2.3")

}
