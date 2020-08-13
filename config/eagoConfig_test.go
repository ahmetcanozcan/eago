package config

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestLoadEagoJSON(t *testing.T) {
	c := qt.New(t)

	jsonC, err := loadEagoJSON(strings.NewReader(`
		{
			"name" : "test",
			"staticPath" : "public"
		}
	`))
	c.Assert(err, qt.IsNil)
	c.Assert(jsonC.Name, qt.Equals, "test")
	c.Assert(jsonC.Port, qt.Equals, 3000)
	c.Assert(jsonC.StaticPath, qt.Equals, "public")
}
