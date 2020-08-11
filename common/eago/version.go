package eago

import "fmt"

const (
	devVersionSuffix   VersionSuffix = "-DEV"
	emptyVersionSuffix VersionSuffix = ""
)

// VersionString :
type VersionString string

// VersionSuffix :
type VersionSuffix string

// Version represents the Eago build version
type Version struct {

	// Major version
	Major int

	// Minor version
	Minor int

	// PatchLevel of Version
	// Increment this for bug fixing
	PatchLevel int

	// Suffix of version
	Suffix VersionSuffix
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d%s",
		v.Major,
		v.Minor,
		v.PatchLevel,
		v.Suffix)
}

// Version presents version as VersionString
func (v Version) Version() VersionString {
	return VersionString(v.String())
}

// Next returns next minor version of this version
func (v Version) Next() Version {
	major := v.Major + int((v.Minor+1)/100)
	minor := (v.Minor + 1) % 100
	return newVersion(major, minor, 0, emptyVersionSuffix)
}

func newVersion(major, minor, patch int, suffix VersionSuffix) Version {
	return Version{
		Major:      major,
		Minor:      minor,
		PatchLevel: patch,
		Suffix:     suffix,
	}
}
