package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	// EagoJSON is the application or module configurations parsed from eago.json file in workspace
	EagoJSON eagoJSON
)

type eagoJSON struct {
	// Path of eago.json file
	path string

	core *eagoJSON
	// Name of the eago appliction or module
	// for modules, name should be repo url
	Name string `json:"name,omitempty"`

	// Version of the application or module
	Version string `json:"version,omitempty"`

	// Author of the application or module
	// it should be formatted as "FULL_NAME <email_address>"
	Author string `json:"author,omitempty"`

	// Module represents this appliction is module or not
	Package bool `json:"package,omitempty"`

	// EagoEnv is target of the application.
	// it would be 'production', 'development' or 'debug'

	EagoEnv string `json:"eagoEnv,omitempty"`

	// Port number that eago application will listen
	Port int `json:"port"`

	// StaticPath
	StaticPath string `json:"staticPath,omitempty"`

	// NotFound
	NotFound string `json:"notFound"`

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

func (ej *eagoJSON) fillDefaults() {
	if ej.NotFound == "" {
		ej.NotFound = "404.html"
	}
	if ej.Port == 0 {
		ej.Port = 3000
	}
	if ej.Dependincies == nil {
		ej.Dependincies = make(map[string]string)
	}
	if ej.DevDependincies == nil {
		ej.DevDependincies = make(map[string]string)
	}
}

func (ej eagoJSON) Save() error {
	data, err := json.MarshalIndent(ej, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(ej.path, data, 0644); err != nil {
		return err
	}

	return nil
}

func parseEago(dirname string) {
	path := filepath.Join(dirname, "./eago.json")
	reader, err := os.Open(path)
	if err != nil {
		// TODO: log warn
		fmt.Println(err)
	}
	config, err := loadEagoJSON(reader)
	if err != nil {
		// TODO: Fatal error
		fmt.Println(err)
		return
	}
	EagoJSON = config
	EagoJSON.path = path
}

func loadEagoJSON(reader io.Reader) (eagoJSON, error) {
	var _json eagoJSON
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return _json, err
	}
	if err := json.Unmarshal(jsonBytes, &_json); err != nil {
		return _json, err
	}
	_json.fillDefaults()
	return _json, nil
}
