package config

import (
	"io"
)

type Installer struct {
	InstallScript string `yaml:"installScript"`
	Executable    string `yaml:"executable"`
}

func (i *Installer) Install(out io.Writer) {
	PlatformInstaller(i, out)
}
