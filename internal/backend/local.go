package backend

import (
	"os"

	c "github.com/aslamcodes/powerappstream-builder/internal/config"
	"github.com/goccy/go-yaml"
)

type LocalBackend struct {
	Location string
}

func (lb *LocalBackend) GetConfig() (*c.Config, error) {
	config, err := os.ReadFile(lb.Location)

	if err != nil {
		return nil, err
	}

	var configData c.Config

	if err := yaml.Unmarshal(config, &configData); err != nil {
		return nil, err
	}

	return &configData, nil
}
