package config

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

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

func (ss *SessionScripts) UpdateSessionScriptConfig(out io.Writer) error {
	if err := os.MkdirAll(filepath.Dir(SESSION_SCRIPT_CONFIG_LOCATION), 0770); err != nil {
		return err
	}

	config, err := json.MarshalIndent(ss, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.Create(SESSION_SCRIPT_CONFIG_LOCATION)

	if err != nil {
		return err
	}

	defer file.Close()

	bw := bufio.NewWriter(file)

	if _, err := bw.Write(config); err != nil {
		return err
	}

	return bw.Flush()
}
