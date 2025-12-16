package config

import (
	"strings"
)

type Image struct {
	Name                    string   `yaml:"name"`
	DisplayName             string   `yaml:"display_name"`
	Description             string   `yaml:"description"`
	EnableDynamicAppCatalog bool     `yaml:"enable_dynamic_app_catalog"`
	UseLatestAgentVersion   bool     `yaml:"use_latest_agent_version"`
	Tags                    []string `yaml:"tags"`
	DryRun                  bool     `yaml:"dry_run"`
}

func (i *Image) Args() []string {
	var args []string

	if i.Name != "" {
		args = append(args, "--name", i.Name)
	}

	if i.DisplayName != "" {
		args = append(args, "--display-name", i.DisplayName)
	}

	if i.Description != "" {
		args = append(args, "--description", i.Description)
	}

	if i.UseLatestAgentVersion {
		args = append(args, "--use-latest-agent-version")
	}

	if i.EnableDynamicAppCatalog {
		args = append(args, "--enable-dynamic-app-catalog")
	}

	if i.DryRun {
		args = append(args, "--dry-run")
	}

	if len(i.Tags) > 0 {
		args = append(args, "--tags")
		for _, tag := range i.Tags {
			parts := strings.SplitSeq(tag, ":")
			for part := range parts {
				args = append(args, part)
			}
		}
	}

	return args
}
