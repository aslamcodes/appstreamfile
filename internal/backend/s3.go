package backend

import (
	c "github.com/aslamcodes/powerappstream-builder/internal/config"
)

type S3Backend struct {
}

func (s3 *S3Backend) GetConfig() (*c.Config, error) {
	return &c.Config{}, nil
}
