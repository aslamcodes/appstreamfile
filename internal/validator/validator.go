package validator

import (
	"fmt"

	"github.com/aslamcodes/powerappstream-builder/internal/config"
)

var ExecPlatformMap = map[string][]string{
	"windows": {"powershell"},
	"unix":    {"bash"},
}

func ValidateConfig(c *config.Config) error {
	validators := []func(*config.Config) error{
		ValidateCatalogApplications,
		ValidateFileDeploys,
		ValidateImage,
		ValidatePaths,
		ValidatePlatforms,
		InstallerValidator,
	}

	for _, v := range validators {
		if err := v(c); err != nil {
			return err
		}
	}

	fmt.Println("Config file validated without any issues")

	return nil
}
