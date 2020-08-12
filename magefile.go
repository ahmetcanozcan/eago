// +build mage

package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/sh"
)

var (
	// default go path. it can be overrided by user using $GOEXE
	goExe = func() string {
		if exe := os.Getenv("GOEXE"); len(exe) > 0 {
			return exe
		}
		return "go"
	}()
	versionSetterPath = func() string {
		dir, _ := os.Getwd()
		return filepath.Join(dir, "common", "eago", "setter.go")
	}()

	ldflags = "-X $PACKAGE/common/eago.commitHash=$COMMIT_HASH -X $PACKAGE/common/eago.buildDate=$BUILD_DATE"
)

const (
	fullName    = "Eago Javascript Runtime"
	packageName = "github.com/ahmetcanozcan/eago"
)

func init() {

}

// Eago builds eago binary
func Eago() error {
	return sh.RunWith(flagEnv(), goExe, "build", "-ldflags", ldflags, ".")
}

// BootstrapFiles evaluates bootstrap files
func BootstrapFiles() {
	sh.Run("node", "scripts/buildBootstrapFiles.js")
}

func flagEnv() map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return map[string]string{
		"FULL_NAME":   fullName,
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
	}
}

var docker = sh.RunCmd("docker")

// Docker builds eago Docker container
func Docker() error {
	if err := docker("build", "-t", "eago", "."); err != nil {
		return err
	}
	docker("rm", "-f", "eago-build")
	if err := docker("run", "--name", "eago-build", "eago ls /go/src/github.com/ahmetcanozcan/eago/"); err != nil {
		return err
	}
	if err := docker("cp", "eago-build:/go/src/github.com/ahmetcanozcan/eago/eago", "."); err != nil {
		return err
	}
	return docker("rm", "eago-build")
}
