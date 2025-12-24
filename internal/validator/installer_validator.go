package validator

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrInvalidExecutableForPlatform = errors.New("Invalid executable for given platform")
)

func InstallerValidator(ctx context.Context, configData *config.Config) error {
	platform := configData.Platform

	platformExecs, exists := ExecPlatformMap[platform]

	if !exists {
		return ErrInvalidPlatform
	}

	for _, e := range configData.Installers {
		if !slices.Contains(platformExecs, e.Executable) {
			return fmt.Errorf("executable %s not supported in platform %s: %w", e.Executable, platform, ErrInvalidExecutableForPlatform)
		}
	}

	return nil
}
