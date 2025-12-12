package service

import (
	"os"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

func TestFileDeploy(t *testing.T) {
	file, err := os.CreateTemp("../../testdata", "test.txt")
	defer func() {
		os.Remove(file.Name())
	}()

	if err != nil {
		t.Errorf("error creating temp file: %v", err)
	}

	f := &config.File{
		Path:    file.Name(),
		Content: "hello world",
	}

	s := FileDeploySvc{}

	s.DeployFile(f)

	content, err := os.ReadFile(f.Path)

	if err != nil {
		t.Errorf("error reading file: %v", err)
	}

	if string(content) != f.Content {
		t.Errorf("expected %s\n, got %s", f.Content, content)
	}

}
