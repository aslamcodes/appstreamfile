package config

import (
	"fmt"
	"io"
	"os/exec"
)

type CatalogConfig struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	DisplayName string `yaml:"display_name"`
	Parameters  string `yaml:"parameters"`
	IconPath    string `yaml:"icon_path"`
	WorkingDir  string `yaml:"working_dir"`
}

func (c CatalogConfig) Args() []string {
	args := []string{}

	if c.Name != "" {
		args = append(args, "--name", c.Name)
	}
	if c.Path != "" {
		args = append(args, "--absolute-app-path", c.Path)
	}
	if c.DisplayName != "" {
		args = append(args, "--display-name", c.DisplayName)
	}
	if c.Parameters != "" {
		args = append(args, "--launch-parameters", c.Parameters)
	}
	if c.IconPath != "" {
		args = append(args, "--absolute-icon-path", c.IconPath)
	}
	if c.WorkingDir != "" {
		args = append(args, "--working-directory", c.WorkingDir)
	}

	return args
}

func (c *CatalogConfig) UpdateStackCatalog(out io.Writer) error {
	_, err := exec.LookPath("image-assistant.exe")

	if err != nil {
		return err
	}

	args := append([]string{"add-application"}, c.Args()...)

	cmd := exec.Command("image-assistant.exe", args...)

	fmt.Println(cmd.String())

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	out.Write(output)

	return nil
}
