package eago

import (
	"fmt"
	"runtime"

	"github.com/ahmetcanozcan/eago/common/types"
)

// Info contains general information about eago software
var Info info = info{
	FullName:   "Eago Javascript Runtime",
	CommitHash: "",
	BuildDate:  "",
}

type info struct {
	// Full name of eago software
	FullName string

	// CommitHash is the current gi revision
	CommitHash string

	// BuildDate is date of current build
	BuildDate string
}

// Version returns version string of current version of eago software
func (i info) Version() VersionString {
	return CurrentVersion.Version()
}

// Header return key value pair of response header
func (i info) Header() types.KeyValueStr {
	return types.NewKeyValueStr("X-Powered-By", "Eago")
}

func (i info) BuildVersionString() string {
	version := CurrentVersion.String()
	if i.CommitHash != "" {
		version += "-" + i.CommitHash
	}
	osArch := runtime.GOOS + "/" + runtime.GOARCH
	return fmt.Sprintf("%s %s %s", i.FullName, version, osArch)
}
