//go:build !windows

package config

import (
	"fmt"
)

func PlatformInstaller(installer *Installer) error {
	fmt.Println("unix in progress 🦔")

	return nil
}
