package config

import (
	"fmt"
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
	args := []string{}

	if i.Name != "" {
		args = append(args, fmt.Sprintf(`--name "%s"`, i.Name))
	}

	if i.DisplayName != "" {
		args = append(args, fmt.Sprintf(`--display-name "%s"`, i.DisplayName))
	}
	if i.Description != "" {
		args = append(args, fmt.Sprintf(`--description "%s"`, i.Description))
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
		tagString := []string{"--tags"}
		for _, tag := range i.Tags {
			parts := strings.SplitSeq(tag, ":")
			for part := range parts {
				tagString = append(tagString, fmt.Sprintf(`"%s"`, part))
			}
		}
		args = append(args, strings.Join(tagString, " "))
	}

	return args
}
