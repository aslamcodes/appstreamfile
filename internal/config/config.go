package config

import "io"

type Config struct {
	Installers []Installer `yaml:"installers"`
}

func (c *Config) Setup(out io.Writer) {
	for _, i := range c.Installers {
		i.Install(out)
	}
}
