package compiler

import (
	"io"
	"io/ioutil"

	babel "github.com/jvatic/goja-babel"
)

// transformToES2015 compile newer version of ecmaSciprt  code to es2015
func transformToES2015(reader io.Reader) (string, error) {
	res, err := babel.Transform(reader, map[string]interface{}{
		"presets": []string{"es2015"},
		"plugins": []string{
			"transform-es2015-block-scoping",
		}})
	if err != nil {
		return "", err
	}
	code, err := ioutil.ReadAll(res)
	if err != nil {
		return "", err
	}
	return string(code), nil

}
