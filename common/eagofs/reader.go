package eagofs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/common/eagrors"
)

// DirReaderHandler handles folder reader results
type DirReaderHandler func(filename string) error

// DirReader reads the folder recursivly
type DirReader struct {
	folderPath string
}

// NewDirReader return new folder reader
func NewDirReader(folderPath string) *DirReader {
	return &DirReader{folderPath}
}

func (r DirReader) Read(handler DirReaderHandler) error {
	return filepath.Walk(
		r.folderPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return eagrors.NewErrorWithCause(err, "Can not read "+path)
			}
			if info.IsDir() {
				return nil
			}
			// eat base dir
			path = path[len(r.folderPath):]
			// change \ to /
			path = strings.Replace(path, "\\", "/", -1)
			return handler(path)
		})
}
