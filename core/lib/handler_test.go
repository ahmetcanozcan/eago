package lib

import (
	"testing"

	"github.com/ahmetcanozcan/eago/common/constants"
	qt "github.com/frankban/quicktest"
)

func TestParseURLFromFilename(t *testing.T) {
	c := qt.New(t)

	testCases := map[string]struct {
		c      string
		method string
	}{
		"index.js":          {c: "", method: constants.HTTP.ALL},
		"test/file.js":      {c: "test/file", method: constants.HTTP.ALL},
		"test/index.js":     {c: "test/", method: constants.HTTP.ALL},
		"test/index.get.js": {c: "test/", method: constants.HTTP.GET},
		"test/file.post.js": {c: "test/file", method: constants.HTTP.POST},
	}
	for key, val := range testCases {
		path, method, err := ParseURLFromFilename(key)
		c.Assert(err, qt.IsNil)
		c.Assert(path, qt.Equals, val.c)
		c.Assert(method, qt.Equals, val.method)
	}
}
