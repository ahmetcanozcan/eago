package core

import "os"

// StartOptions :
type StartOptions struct {
	// Directory that application located
	// By default it's Current Working Directory
	AppDir string
}

func (opt *StartOptions) fillDefaults() {
	cwd, _ := os.Getwd()
	if len(opt.AppDir) == 0 {
		opt.AppDir = cwd
	}
}
