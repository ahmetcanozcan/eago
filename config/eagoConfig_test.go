package config

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/spf13/viper"
)

func TestLoadEagoJSON(t *testing.T) {
	c := qt.New(t)
	v := viper.New()
	v.Set("name", testName)
	jsonC, err := loadEagoJSON(v)
	c.Assert(err, qt.IsNil)
	c.Assert(jsonC.Name, qt.Equals, testName)
}
