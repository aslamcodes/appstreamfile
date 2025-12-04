package config

import (
	"fmt"
	"io"
)

type Installer struct {
	InstallScript string `yaml:"installScript"`
	Executable    string `yaml:"executable"`
}

type Config struct {
	Installers []Installer `yaml:"installers"`
}

func (i *Installer) Install(out io.Writer) {
	fmt.Fprintln(out, fmt.Sprintf("%s %s", i.Executable, i.InstallScript))
}
