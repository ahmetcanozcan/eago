package text

import "fmt"

// Position is a source position of a text file or stream
type Position struct {
	Filename     string
	Offset       int
	LineNumber   int
	ColumnNumber int
}

func (p Position) String() string {
	return fmt.Sprintf("%s %d:%d", p.Filename, p.LineNumber, p.ColumnNumber)
}
