package config

type Image struct {
    Name                    string   `yaml:"name"`
    DisplayName             string   `yaml:"display_name"`
    Description             string   `yaml:"description"`
    EnableDynamicAppCatalog bool     `yaml:"enable_dynamic_app_catalog"`
    UseLatestAgentVersion   bool     `yaml:"use_latest_agent_version"`
    Tags                    []string `yaml:"tags"`
    DryRun                  bool     `yaml:"dry_run"`
}
