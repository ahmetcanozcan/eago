package compiler

import (
	"os"

	"github.com/ahmetcanozcan/eago/common/eagrors"
)

// Script :
type Script struct {
	Filename string
	Code     string
}

// ReadScript reads file given filename and store it
func ReadScript(filename string) (*Script, error) {
	erStr := "Can not read script " + filename
	file, err := os.Open(filename)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, erStr)
	}
	code, err := transformToES2015(file)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, erStr)
	}
	return &Script{filename, code}, nil
}
