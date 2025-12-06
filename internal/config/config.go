package config

import (
	"fmt"
	"io"
)

type Config struct {
	Installers     []Installer     `yaml:"installers"`
	Files          []File          `yaml:"files"`
	Catalogs       []CatalogConfig `yaml:"catalog"`
	SessionScripts SessionScripts  `yaml:"session_scripts"`
}

func (c *Config) Setup(out io.Writer) error {
	for _, f := range c.Files {
		err := f.Deploy(out)

		if err != nil {
			return fmt.Errorf("error deploying file %s: %w", f.Path, err)
		}
	}

	for _, i := range c.Installers {
		err := i.Install(out)

		if err != nil {
			return fmt.Errorf("error installing %s: %w", i.Executable+i.InstallScript, err)
		}
	}

	if err := c.SessionScripts.UpdateSessionScriptConfig(out); err != nil {
		return fmt.Errorf("error configuring session scripts: %w", err)
	}

	for _, catalog := range c.Catalogs {
		if err := catalog.UpdateStackCatalog(out); err != nil {
			return fmt.Errorf("error updating stack catalog: %w", err)
		}
	}

	return nil
}
