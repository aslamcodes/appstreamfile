package backend

import (
	"context"
	"fmt"
	"os"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/goccy/go-yaml"
)

type LocalBackend struct {
	Location string
}

func (lb *LocalBackend) GetConfig(ctx context.Context) (*config.Config, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Attempting to fetch config from local backend at %s\n", lb.Location)

	data, err := os.ReadFile(lb.Location)

	if err != nil {
		return nil, fmt.Errorf("failed to read from location %s: %w", lb.Location, err)
	}

	var configData config.Config

	if err := yaml.Unmarshal(data, &configData); err != nil {
		return nil, fmt.Errorf("failed to parse config data, config data or formatting is invalid: %w", err)
	}

	fmt.Printf("Builder has successfully parsed the config file from backend\n")

	return &configData, nil
}

func NewLocalBackend(location string) (BackendSource, error) {
	return &LocalBackend{
		Location: location,
	}, nil
}
