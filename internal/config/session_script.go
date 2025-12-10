package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func SessionScriptLocation() string {
	if runtime.GOOS == "windows" {
		return `C:\Appstream\SessionScripts\config.json`
	}
	return "/opt/appstream/SessionScripts/config.json"
}

func (ss *SessionScripts) UpdateSessionScriptConfig(location string) error {
	fmt.Println("Configuring session scripts")

	if err := os.MkdirAll(filepath.Dir(location), 0770); err != nil {
		return err
	}

	config, err := json.MarshalIndent(ss, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.Create(location)

	if err != nil {
		return err
	}

	fmt.Println("Sessions scripts are configured successfully at", location)

	defer file.Close()

	bw := bufio.NewWriter(file)

	if _, err := bw.Write(config); err != nil {
		return err
	}

	return bw.Flush()
}
