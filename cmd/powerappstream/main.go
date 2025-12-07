package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aslamcodes/appstream-file/internal/backend"
	"github.com/aslamcodes/appstream-file/internal/logger"
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
	switch sourceType {
	case "local":
		backend := backend.LocalBackend{
			Location: location,
		}

		config, err := backend.GetConfig()

		if err != nil {
			return fmt.Errorf("failed to fetch config from backend: %w", err)
		}

		err = config.Setup()

		if err != nil {
			return fmt.Errorf("error setting up config: %w", err)
		}

	default:
		return fmt.Errorf("invalid source provided")
	}

	return nil
}
