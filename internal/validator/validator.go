package validator

import (
	"context"
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var ExecPlatformMap = map[string][]string{
	"windows": {"powershell"},
	"unix":    {"bash"},
}

func ValidateConfig(ctx context.Context, c *config.Config) error {
	validators := []func(context.Context, *config.Config) error{
		ValidateCatalogApplications,
		ValidateFileDeploys,
		ValidateImage,
		ValidatePaths,
		ValidatePlatforms,
		InstallerValidator,
	}

	for _, v := range validators {
		if err := v(ctx, c); err != nil {
			return err
		}
	}

	fmt.Println("Config file validated without any issues")

	return nil
}
