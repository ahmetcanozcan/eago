package eagrors

import "github.com/ahmetcanozcan/eago/common/text"

// FileError represent errors that invoked when handling, executing, parsing a file
type FileError interface {
	error
	Position() text.Position
}
