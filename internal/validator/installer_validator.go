package validator

import (
	"errors"
	"fmt"
	"slices"

	"github.com/aslamcodes/powerappstream-builder/internal/config"
)

var (
	ErrInvalidExecutableForPlatform = errors.New("Invalid executable for given platform")
	ErrInvalidPlatform              = errors.New("Platform not supported")
)

func InstallerValidator(configData *config.Config) error {
	execPlatformMap := map[string][]string{
		"windows": {"powershell"},
		"unix":    {"bash"},
	}

	platform := configData.Platform

	platformExecs, exists := execPlatformMap[platform]

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
