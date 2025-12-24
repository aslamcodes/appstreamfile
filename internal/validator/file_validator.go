package validator

import (
	"context"
	"errors"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrFileDeployPathMissing = errors.New("file deploy path should not be null")
)

func ValidateFileDeploys(ctx context.Context, c *config.Config) error {
	for _, file := range c.Files {
		if file.Path == "" {
			return ErrFileDeployPathMissing
		}
	}

	return nil
}
