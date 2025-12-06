package backend

import (
	"fmt"
	"os"

	"github.com/aslamcodes/powerappstream-builder/internal/config"
	"github.com/goccy/go-yaml"
)

type LocalBackend struct {
	Location string
}

func (lb *LocalBackend) GetConfig() (*config.Config, error) {
	data, err := os.ReadFile(lb.Location)

	if err != nil {
		return nil, fmt.Errorf("failed to read from location %s: %w", lb.Location, err)
	}

	var configData config.Config

	if err := yaml.Unmarshal(data, &configData); err != nil {
		return nil, fmt.Errorf("failed to parse config data, config data or formatting is invalid: %w", err)
	}

	return &configData, nil
}
