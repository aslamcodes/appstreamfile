package config

type baseConfig struct {
	Filename     string `yaml:"filename"`
	Arguments    string `yaml:"arguments"`
	S3LogEnabled bool   `yaml:"s3LogEnabled"`
}

type Executables struct {
	SystemContext baseConfig `yaml:"system_context"`
	UserContext   baseConfig `yaml:"user_context"`
}

type SessionScript struct {
	SessionStart       Executables `yaml:"session_start"`
	SessionTermination Executables `yaml:"session_termination"`
}
