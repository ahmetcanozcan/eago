package eagofs

import "os"

// MkdirAll create directory
func MkdirAll(dirpath string, umask os.FileMode) error {
	return os.MkdirAll(dirpath, umask)
}
