package config

import "io"

type Config struct {
	Installers     []Installer    `yaml:"installers"`
	Files          []File         `yaml:"files"`
	SessionScripts SessionScripts `yaml:"session_scripts"`
}

func (c *Config) Setup(out io.Writer) error {
	for _, f := range c.Files {
		err := f.Deploy(out)

		if err != nil {
			return err
		}
	}

	for _, i := range c.Installers {
		err := i.Install(out)

		if err != nil {
			return err
		}
	}

	if err := c.SessionScripts.UpdateSessionScriptConfig(out); err != nil {
		return err
	}

	return nil
}
