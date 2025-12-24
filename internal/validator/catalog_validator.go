package validator

import (
	"context"
	"errors"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

var (
	ErrEmptyCatalogName     = errors.New("catalog application name cannot be empty")
	ErrEmptyCatalogPath     = errors.New("catalog application path cannot be empty")
	ErrEmptyCatalogNamePath = errors.New("catalog application path and name cannot be empty")
)

func ValidateCatalogApplications(ctx context.Context, configData *config.Config) error {
	for _, c := range configData.Catalogs {
		if c.Name == "" && c.Path == "" {
			return ErrEmptyCatalogNamePath
		}
		if c.Name == "" {
			return ErrEmptyCatalogName
		}

		if c.Path == "" {
			return ErrEmptyCatalogPath
		}
	}

	return nil
}
