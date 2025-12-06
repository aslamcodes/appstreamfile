package backend

import (
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
		return nil, err
	}

	var configData config.Config

	if err := yaml.Unmarshal(data, &configData); err != nil {
		return nil, err
	}

	return &configData, nil
}
