package config

type Executable struct {
	Context      string `yaml:"context" json:"context"`
	Filename     string `yaml:"filename" json:"filename"`
	Arguments    string `yaml:"arguments" json:"arguments"`
	S3LogEnabled bool   `yaml:"s3LogEnabled" json:"s3LogEnabled"`
}

type SessionConfig struct {
	Executables []Executable `yaml:"executables" json:"executables"`
	WaitingTime int          `yaml:"waitingTime" json:"waitingTime"`
}

type SessionScripts struct {
	SessionStart       SessionConfig `yaml:"session_start" json:"SessionStart"`
	SessionTermination SessionConfig `yaml:"session_termination" json:"SessionTermination"`
}
