package validator

import (
	"context"
	"errors"
	"slices"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrPlatformMissing = errors.New("config file does not contain platform specification")
	ErrInvalidPlatform = errors.New("Platform not supported")
)

func ValidatePlatforms(ctx context.Context, c *config.Config) error {
	if c.Platform == "" {
		return ErrPlatformMissing
	}

	supportedPlatforms := make([]string, len(ExecPlatformMap))

	for k := range ExecPlatformMap {
		supportedPlatforms = append(supportedPlatforms, k)
	}

	if !slices.Contains(supportedPlatforms, c.Platform) {
		return ErrInvalidPlatform
	}

	return nil
}
