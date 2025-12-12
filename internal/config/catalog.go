package config

import (
	"fmt"
	"strings"
)

type CatalogConfig struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	DisplayName string `yaml:"display_name"`
	Parameters  string `yaml:"parameters"`
	IconPath    string `yaml:"icon_path"`
	WorkingDir  string `yaml:"working_dir"`
}

func (c *CatalogConfig) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("name - %s\n", c.Name))
	builder.WriteString(fmt.Sprintf("absolute-app-path - %s\n", c.Path))
	builder.WriteString(fmt.Sprintf("display-name - %s\n", c.DisplayName))
	builder.WriteString(fmt.Sprintf("launch-parameters - %s\n", c.Parameters))
	builder.WriteString(fmt.Sprintf("absolute-icon-path - %s\n", c.IconPath))
	builder.WriteString(fmt.Sprintf("working-directory - %s\n", c.WorkingDir))

	return builder.String()
}

func (c *CatalogConfig) Args() []string {
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
