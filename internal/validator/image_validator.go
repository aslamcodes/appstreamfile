package validator

import (
	"errors"

	"github.com/aslamcodes/powerappstream-builder/internal/config"
)

var (
	ErrInvalidParametersCreateImage = errors.New("invalid parameters for create image operation")
)

func ValidateImage(c *config.Config) error {
	if c.Image.Name == "" {
		return ErrInvalidParametersCreateImage
	}
	return nil
}
