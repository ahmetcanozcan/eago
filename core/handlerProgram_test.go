package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ahmetcanozcan/eago/common/constants"
	qt "github.com/frankban/quicktest"
)

func TestParseHandlers(t *testing.T) {
	c := qt.New(t)
	cwd, _ := os.Getwd()
	dirname := filepath.Join(cwd, "..", "test_files", "handlers")
	fmt.Println(dirname)
	hb, err := parseHandlers(dirname)
	c.Assert(err, qt.IsNil)
	_ = map[string]bool{
		"file/test": true, // ALL
		"test":      true, // ALL
		"another":   true, // POST
		"file":      true, // GET
	}

	testURLs := []struct {
		method string
		url    string
	}{
		{method: constants.HTTP.GET, url: "/file/"},
		{method: constants.HTTP.GET, url: "/file/"},
		{method: constants.HTTP.GET, url: "file/test"},
		{method: constants.HTTP.POST, url: "file/test"},
		{method: constants.HTTP.ALL, url: "test"},
	}

	for _, v := range testURLs {
		flag := false
		for _, bundle := range hb {
			checked := bundle.URLPath.Check(v.url)
			if checked {
				prog, err := bundle.getProgram(v.method)
				c.Assert(err, qt.IsNil)
				c.Assert(prog, qt.Not(qt.IsNil))
				flag = true
				break
			}
		}
		c.Assert(flag, qt.IsTrue)
	}
}
