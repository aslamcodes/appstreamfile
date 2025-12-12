package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

type SessionScriptSvc struct{}

func SessionScriptLocation() string {
	if runtime.GOOS == "windows" {
		return `C:\Appstream\SessionScripts\config.json`
	}
	return "/opt/appstream/SessionScripts/config.json"
}

func (svc *SessionScriptSvc) UpdateSessionScriptConfig(location string, ss config.SessionScripts) error {
	fmt.Println("Configuring session scripts")

	if err := os.MkdirAll(filepath.Dir(location), 0770); err != nil {
		return err
	}

	config, err := json.MarshalIndent(ss, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.Create(location)

	if err != nil {
		return err
	}

	fmt.Println("Sessions scripts are configured successfully at", location)

	defer file.Close()

	bw := bufio.NewWriter(file)

	if _, err := bw.Write(config); err != nil {
		return err
	}

	return bw.Flush()
}
