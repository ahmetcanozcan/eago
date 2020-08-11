package config

import (
	"fmt"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

const testName = "testName"

func TestReadConfigJSON(t *testing.T) {
	c := qt.New(t)

	reader := strings.NewReader(fmt.Sprintf(`
	{
		"name": "%s"
	}
	`, testName))
	v := readConfig(reader, map[string]interface{}{
		"port":   300,
		"module": true,
	}, "json")
	c.Assert(v.GetString("name"), qt.Equals, testName)
	c.Assert(v.GetInt("port"), qt.Equals, 300)
	c.Assert(v.GetBool("module"), qt.Equals, true)
}
