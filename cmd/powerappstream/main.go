package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/aslamcodes/powerappstream-builder/internal/backend"
)

func main() {
	source := flag.String("source", "", "The source to pick actions from")
	location := flag.String("location", "", "The config file location")

	flag.Parse()

	if err := run(*source, *location, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(sourceType string, location string, out io.Writer) error {
	switch sourceType {
	case "local":
		backend := backend.LocalBackend{
			Location: location,
		}

		config, err := backend.GetConfig()

		if err != nil {
			return err
		}

		config.Setup(out)

		// case "powerappstream":
		// case "s3":
		// case "git":

	default:
		return fmt.Errorf("Invalid source provided")
	}

	return nil
}
