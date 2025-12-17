package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aslamcodes/appstreamfile/internal/backend"
	"github.com/aslamcodes/appstreamfile/internal/logger"
	"github.com/aslamcodes/appstreamfile/internal/service"
	"github.com/aslamcodes/appstreamfile/internal/validator"
)

func main() {
	source := flag.String("source", "", "The source to pick actions from")
	location := flag.String("location", "", "The config file location")

	flag.Parse()

	logger.Init()

	if err := run(*source, *location); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(sourceType string, location string) error {
	var backendSource backend.BackendSource
	var err error

	switch sourceType {
	case "local":
		backendSource, err = backend.NewLocalBackend(location)
	case "s3":
		backendSource, err = backend.NewS3Backend("mymediaiso", "state/config_win.yaml", "", "appstream_machine_role")
	default:
		return fmt.Errorf("invalid source provided")
	}

	if err != nil {
		return fmt.Errorf("unable to create backend source: %w", err)
	}

	config, err := backendSource.GetConfig()

	if err != nil {
		return fmt.Errorf("failed to fetch config from backend: %w", err)
	}

	if err := validator.ValidateConfig(config); err != nil {
		return fmt.Errorf("config file validation failed: %w", err)
	}

	err = service.ImplementConfig(config)

	if err != nil {
		return fmt.Errorf("error setting up config: %w", err)
	}

	return nil
}
