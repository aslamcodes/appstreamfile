package backend

import (
	"context"
	"fmt"
	"io"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/goccy/go-yaml"
)

type S3Backend struct {
	Bucket    string
	Key       string
	VersionId string
	Client    S3Client
}

func (s3Backend *S3Backend) GetConfig(ctx context.Context) (*config.Config, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if s3Backend.Client == nil {
		return nil, fmt.Errorf("client is nil")
	}

	out, err := s3Backend.Client.GetObject(ctx, s3Backend.Bucket, s3Backend.Key, s3Backend.VersionId)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch object from s3: %w", err)
	}

	defer out.Body.Close()

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

func NewS3Backend(ctx context.Context, bucket, key, versionId, profile string) (BackendSource, error) {
	client, err := NewS3Client(ctx, profile)

	if err != nil {
		return nil, fmt.Errorf("not able to create s3 client: %w", err)
	}

	return &S3Backend{
		Bucket:    bucket,
		Key:       key,
		VersionId: versionId,
		Client:    client,
	}, nil
}
