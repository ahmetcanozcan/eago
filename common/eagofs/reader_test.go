package eagofs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/frankban/quicktest"
)

func TestDirReader(t *testing.T) {
	c := quicktest.New(t)

	cwd, _ := os.Getwd()
	dirname := filepath.Join(cwd, "..", "..", "common")
	reader := NewDirReader(dirname)
	flag := false
	reader.Read(func(filename string) error {
		flag = flag || filename == "/eagofs/reader_test.go"
		return nil
	})

	c.Assert(flag, quicktest.Equals, true)

}
