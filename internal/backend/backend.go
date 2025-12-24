package backend

import (
	"context"

	c "github.com/aslamcodes/appstreamfile/internal/config"
)

type BackendSource interface {
	GetConfig(ctx context.Context) (*c.Config, error)
}
