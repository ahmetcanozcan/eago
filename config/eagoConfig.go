package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	// EagoJSON is the application or module configurations parsed from eago.json file in workspace
	EagoJSON eagoJSON
)

type eagoJSON struct {

	// Name of the eago appliction or module
	// for modules, name should be repo url
	Name string `json:"name"`

	// Version of the application or module
	Version string `json:"version"`

	// Author of the application or module
	// it should be formatted as "FULL_NAME <email_address>"
	Author string `json:"author"`

	// Module represents this appliction is module or not
	Module bool `json:"module"`

	// EagoEnv is target of the application.
	// it would be 'production', 'development' or 'debug'
	EagoEnv string `json:"eagoEnv"`

	// Port number that eago application will listen
	Port int `json:"port"`

	// StaticPath
	StaticPath string `json:"staticPath"`

	Dependincies    map[string]string `json:"dependincies"`
	DevDependincies map[string]string `json:"devDependincies"`
}

// IsProduction returns application executed for production or not
func (ej eagoJSON) IsProduction() bool {
	return ej.EagoEnv == "production"
}

// IsDevelopment returns application executed for development or not
func (ej eagoJSON) IsDevelopment() bool {
	return ej.EagoEnv == "development"
}

func (ej eagoJSON) Address() string {
	return fmt.Sprintf(":%d", ej.Port)
}

var eagoJSONDefaultValues = map[string]interface{}{
	"port":       3000,
	"name":       "unnamed",
	"version":    "1.0.0",
	"eagoEnv":    "development",
	"staticPath": ".",
}

func parseEago(dirname string) {
	reader, err := os.Open(filepath.Join(dirname, "./eago.json"))
	if err != nil {
		// TODO: log warn
		fmt.Println(err)
	}
	v := readConfig(reader, eagoJSONDefaultValues, "json")
	config, err := loadEagoJSON(v)
	if err != nil {
		// TODO: Fatal error
		fmt.Println(err)
		return
	}
	EagoJSON = config
}

func loadEagoJSON(cfg *viper.Viper) (eagoJSON, error) {
	var json eagoJSON
	if err := cfg.Unmarshal(&json); err != nil {
		return json, err
	}
	return json, nil
}
