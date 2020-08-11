package modules

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestReadFileLineByLine(t *testing.T) {
	c := qt.New(t)
	testStr := []string{
		"This is",
		"multiline text message",
		"for testing",
		"readfile func",
	}

	res := make([]string, len(testStr))
	i := 0
	readFileLineByLine(strings.NewReader(strings.Join(testStr, "\n")), func(line string) {
		res[i] = line
		i++
	})

	c.Assert(res, qt.CmpEquals(), testStr)

}
