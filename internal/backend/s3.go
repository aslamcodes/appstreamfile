package backend

import (
	"context"
	"fmt"
	"io"

	"github.com/aslamcodes/appstreamfile/internal/config"
	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/goccy/go-yaml"
)

type S3Backend struct {
	Bucket    string
	Key       string
	Region    string
	Profile   string
	Endpoint  string
	VersionID string
}

func (s3Backend *S3Backend) GetConfig() (*config.Config, error) {
	opts := []func(*s3Config.LoadOptions) error{}

	if s3Backend.Profile != "" {
		opts = append(opts, s3Config.WithSharedConfigProfile(s3Backend.Profile))
	}

	ctx := context.Background()

	cfg, err := s3Config.LoadDefaultConfig(ctx, opts...)

	if err != nil {
		return nil, fmt.Errorf("failed to load SDK configuration, %w", err)
	}

	client := s3.NewFromConfig(cfg)

	objectConfig := s3.GetObjectInput{
		Bucket: &s3Backend.Bucket,
		Key:    &s3Backend.Key,
	}

	if s3Backend.VersionID != "" {
		objectConfig.VersionId = &s3Backend.VersionID
	}

	out, err := client.GetObject(ctx, &objectConfig)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch object from s3, %w: ", err)
	}

	content, err := io.ReadAll(out.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var configData config.Config

	if err := yaml.Unmarshal(content, &configData); err != nil {
		return nil, fmt.Errorf("failed to parse config data, config data or formatting is invalid: %w", err)
	}

	fmt.Printf("Builder has successfully parsed the config file from backend\n")

	return &configData, nil
}
