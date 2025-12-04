package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/aslamcodes/powerappstream-builder/internal/agent"
	"github.com/aslamcodes/powerappstream-builder/internal/backend"
)

func main() {
	source := flag.String("source", "", "The source to pick actions from")
	location := flag.String("location", "", "The config file location")

	flag.Parse()

	run(*source, *location, os.Stdout)
}

func run(sourceType string, location string, out io.Writer) {
	agent := agent.Agent{}
	switch sourceType {
	case "local":
		backend := backend.LocalBackend{
			Location: location,
		}

		config, err := backend.GetConfig()

		if err != nil {
			fmt.Fprintln(out, err)
		}

		agent.HandleConfig(config, os.Stdout)
	case "powerappstream":
	case "s3":
	case "git":
	case "consul":
	case "http":
	case "k8":
	case "azurerm":
	default:
	}
}
