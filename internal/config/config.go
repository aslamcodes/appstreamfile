package config

import "io"

type Config struct {
	Installers []Installer `yaml:"installers"`
	Files []File `yaml:"files"`
}

func (c *Config) Setup(out io.Writer) error {
	for _, i := range c.Installers {
		err := i.Install(out)

		if err != nil {
			return err
		}
	}

	return nil
}
