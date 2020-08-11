package eagofs

import (
	"io"
	"os"
)

// IsExist checks for a folder or file is exist or not
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// IsDir checks for given path  is a folder or not
// if file not exist then returns false
func IsDir(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

// IsEmpty checks a folder is empty or not
func IsEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}
