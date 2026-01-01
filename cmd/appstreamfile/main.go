package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/aslamcodes/appstreamfile/internal/backend"
	"github.com/aslamcodes/appstreamfile/internal/logger"
	"github.com/aslamcodes/appstreamfile/internal/service"
	"github.com/aslamcodes/appstreamfile/internal/validator"
)

type RunOptions struct {
	location   string
	SourceType string
	bucket     string
	key        string
	versionId  string
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	source := flag.String("source", "", "Configuration source: s3 or local")
	location := flag.String("location", "", "Local filesystem path to the config file")
	bucket := flag.String("bucket", "", "S3 bucket containing the config file")
	key := flag.String("key", "", "S3 object key for the config file")
	versionId := flag.String("version-id", "", "Optional S3 object version ID")

	flag.Parse()

	logger.Init()

	runOptions := &RunOptions{
		SourceType: *source,
		location:   *location,
		bucket:     *bucket,
		key:        *key,
		versionId:  *versionId,
	}

	if err := run(ctx, runOptions); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, opts *RunOptions) error {
	var backendSource backend.BackendSource
	var err error

	switch opts.SourceType {
	case "local":
		if opts.location == "" {
			return fmt.Errorf("location of config file must be provided")
		}
		backendSource, err = backend.NewLocalBackend(opts.location)
	case "s3":
		if opts.bucket == "" || opts.key == "" {
			return fmt.Errorf("missing required S3 options: bucket and key")
		}
		backendSource, err = backend.NewS3Backend(ctx, opts.bucket, opts.key, opts.versionId, "appstream_machine_role")

	default:
		return fmt.Errorf("invalid source provided")
	}

	if err != nil {
		return fmt.Errorf("unable to create backend source: %w", err)
	}

	config, err := backendSource.GetConfig(ctx)

	if err != nil {
		return fmt.Errorf("failed to fetch config from backend: %w", err)
	}

	if err := validator.ValidateConfig(ctx, config); err != nil {
		return fmt.Errorf("config file validation failed: %w", err)
	}

	err = service.ImplementConfig(ctx, config)

	if err != nil {
		return fmt.Errorf("error setting up config: %w", err)
	}

	return nil
}
