package validator

import (
	"errors"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrFileDeployPathMissing = errors.New("file deploy path should not be null")
)

func ValidateFileDeploys(c *config.Config) error {
	for _, file := range c.Files {
		if file.Path == "" {
			return ErrFileDeployPathMissing
		}
	}

	return nil
}
