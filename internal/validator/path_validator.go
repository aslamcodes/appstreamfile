package validator

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrInvalidPathForUnix    = errors.New("path is invalid for unix")
	ErrInvalidPathForWindows = errors.New("path is invalid for windows")
)

func ValidatePaths(ctx context.Context, c *config.Config) error {
	var (
		winDrive = regexp.MustCompile(`^[A-Za-z]:`)
		uncPath  = regexp.MustCompile(`^\\\\`)
	)

	var validatePath = func(path string) error {
		switch c.Platform {
		case "windows":
			if strings.HasPrefix(path, "/") && !strings.HasPrefix(path, "\\\\") {
				return fmt.Errorf("path (%s) is invalid: %w", path, ErrInvalidPathForWindows)
			}

		case "unix":
			if winDrive.MatchString(path) || uncPath.MatchString(path) {
				return fmt.Errorf("path (%s) is invalid: %w", path, ErrInvalidPathForUnix)
			}
		}

		return nil
	}

	for _, p := range c.Files {
		if err := validatePath(p.Path); err != nil {
			return err
		}
	}

	for _, catalog := range c.Catalogs {
		if err := validatePath(catalog.Path); err != nil {
			return err
		}

		if err := validatePath(catalog.IconPath); err != nil {
			return err
		}
	}

	return nil
}
