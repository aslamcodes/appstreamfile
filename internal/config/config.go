package config

import (
	"fmt"
)

type Config struct {
	Platform       string          `yaml:"platform"`
	Installers     []Installer     `yaml:"installers"`
	Files          []File          `yaml:"files"`
	Catalogs       []CatalogConfig `yaml:"catalog"`
	SessionScripts SessionScripts  `yaml:"session_scripts"`
}

func (c *Config) Setup() error {
	if err := c.SessionScripts.UpdateSessionScriptConfig(); err != nil {
		return fmt.Errorf("error configuring session scripts: %w", err)
	}

	for _, f := range c.Files {
		fmt.Println("Deploying file", f.Path)
		err := f.Deploy()

		if err != nil {
			return fmt.Errorf("error deploying file %s: %w", f.Path, err)
		}
	}

	for _, i := range c.Installers {
		fmt.Println("Executing installer with", i.Executable)
		err := i.Install()

		if err != nil {
			return fmt.Errorf("error installing %s: %w", i.Executable+i.InstallScript, err)
		}
	}

	for _, catalog := range c.Catalogs {
		if err := catalog.UpdateStackCatalog(); err != nil {
			return fmt.Errorf("error updating stack catalog: %w", err)
		}
	}

	return nil
}
