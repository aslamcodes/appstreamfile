//go:build windows

package config

import (
	"fmt"
	"io"
)

func PlatformInstaller(installer *Installer, out io.Writer) {
	fmt.Fprintf(out, "WIP for Windows 🪟")
}
