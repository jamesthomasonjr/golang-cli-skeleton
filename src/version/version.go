package version

import (
	"fmt"
    "os"
    "path/filepath"
)

var name string
var short string
var long string
var version string
var gitVersion string
var goVersion string
var buildDate string
var buildOS string
var buildArch string
var defaultURI string

func ExecutableName() (string, error) {
    executable, err := os.Executable()
    if err == nil {
		executable = fmt.Sprintf("%s", filepath.Base(executable))
	}

	return executable, err
}

func ApplicationName() string {
	return name
}

func ShortDescription() string {
	return short
}

func LongDescription() string {
	return long
}

func MainVersionLine() string {
	return fmt.Sprintf("%s version %s, built %s", name, version, buildDate)
}

func GitVersionLine() string {
	return fmt.Sprintf("git version %s", gitVersion)
}

func BuildVersionLine() string {
	return fmt.Sprintf("built on %s/%s using go version %s", buildOS, buildArch, goVersion)
}
