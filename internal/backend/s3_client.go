package backend

import (
	"context"
	"fmt"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	GetObject(ctx context.Context, bucket string, key string, versionId string) (*s3.GetObjectOutput, error)
}

type S3BackendClient struct {
	s3Client *s3.Client
}

func (c *S3BackendClient) GetObject(ctx context.Context, bucket string, key string, versionId string) (*s3.GetObjectOutput, error) {
	objectInput := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	if versionId != "" {
		objectInput.VersionId = &versionId
	}

	return c.s3Client.GetObject(ctx, objectInput)
}

func NewS3Client(ctx context.Context, profile string) (S3Client, error) {
	opts := []func(*s3Config.LoadOptions) error{}

	if profile != "" {
		opts = append(opts, s3Config.WithSharedConfigProfile(profile))
	}

	cfg, err := s3Config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load SDK configuration: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	backendClient := S3BackendClient{
		s3Client: client,
	}

	return &backendClient, nil
}
