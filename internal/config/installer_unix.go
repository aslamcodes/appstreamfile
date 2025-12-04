//go:build !windows

package config

import (
	"fmt"
	"io"
)

func PlatformInstaller(installer *Installer, out io.Writer) error {
	fmt.Fprintf(out, "unix in progress 🦔")

	return nil
}
