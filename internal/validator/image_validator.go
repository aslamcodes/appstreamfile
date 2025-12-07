package validator

import (
	"errors"
	"strings"

	"github.com/aslamcodes/powerappstream-builder/internal/config"
)

var (
	ErrInvalidParametersCreateImage = errors.New("invalid parameters for create image operation")
	ErrInvalidTagsCreateImage = errors.New("format invalid for create-image tags (key1:value1)")
)

func ValidateImage(c *config.Config) error {
	if c.Image.Name == "" {
		return ErrInvalidParametersCreateImage
	}

	for _, tag := range c.Image.Tags {
		if !strings.ContainsRune(tag, ':') {
			return ErrInvalidTagsCreateImage
		}
	}
	return nil
}
