package config_test

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

func Test(t *testing.T) {
	dir := os.TempDir()
	config_file_path := path.Join(dir, "config.json")

	ss := &config.SessionScripts{
		SessionStart: config.SessionConfig{
			Executables: []config.Executable{
				{
					Context:      "system",
					Filename:     "sample.exe",
					Arguments:    "-h",
					S3LogEnabled: false,
				},
				{
					Context:      "user",
					Filename:     "sample.exe",
					Arguments:    "-h",
					S3LogEnabled: false,
				},
			},
			WaitingTime: 0,
		},
		SessionTermination: config.SessionConfig{
			Executables: []config.Executable{
				{
					Context:      "user",
					Filename:     "sample.exe",
					Arguments:    "-h",
					S3LogEnabled: false,
				},
				{
					Context:      "system",
					Filename:     "sample.exe",
					Arguments:    "-h",
					S3LogEnabled: false,
				},
			},
			WaitingTime: 0,
		},
	}

	ss.UpdateSessionScriptConfig(config_file_path)

	content, err := os.ReadFile(config_file_path)

	if err != nil {
		t.Errorf("failed to read config content: %v", err)
	}

	var actual *config.SessionScripts = &config.SessionScripts{}

	if err := json.Unmarshal(content, actual); err != nil {
		t.Errorf("failed to unmarshal content: %v", err)
	}

	if !reflect.DeepEqual(actual, ss) {
		t.Fatalf("config mismatch\nwant: %v\ngot: %v", ss, actual)
	}
}
