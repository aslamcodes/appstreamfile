package config

type BaseConfig struct {
	Filename     string `yaml:"filename"`
	Arguments    string `yaml:"arguments"`
	S3LogEnabled bool   `yaml:"s3LogEnabled"`
}

type Executables struct {
	SystemContext BaseConfig `yaml:"system_context"`
	UserContext   BaseConfig `yaml:"user_context"`
}

type SessionScripts struct {
	SessionStart       Executables `yaml:"session_start"`
	SessionTermination Executables `yaml:"session_termination"`
}
