package lib

import (
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/common/constants"
	"github.com/ahmetcanozcan/eago/common/eagrors"
)

// ParseURLFromFilename take file name of script.
// returns URL and method if it's defined in filename,
//  otherwise returns ALL method
// i.e    "/test/file.js" -> ("/test/file", "ALL")
//				"/test/file.post.js" -> ("test/file", "POST")
//				"/test/folder/index.js" ("test/folder","ALL")
func ParseURLFromFilename(filename string) (string, string, error) {
	dirname, filename := filepath.Split(filename)
	fileParts := strings.Split(filename, ".")
	method := constants.HTTP.ALL

	// If length of fileparts is 2,
	// filename does not contain any information about method
	if len(fileParts) == 2 || len(fileParts) == 3 {
		// if filename is not index
		// add filename to path
		if fileParts[0] != "index" {
			dirname += fileParts[0]
		}
		// if length of fileparts is 3
		// filename have information about method
		if len(fileParts) == 3 {
			method = strings.ToUpper(fileParts[1])
		}
	} else {
		// Otherwise filePath is invalid so invoke an error
		return "", "", eagrors.NewError("Invalid filename in " + filename)
	}

	return dirname, method, nil
}
