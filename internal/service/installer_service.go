package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

type InstallerSvc struct {
}

func (s *InstallerSvc) Install(inst *config.Installer) error {
	runScript := func(exe, ext, body string, argsPrefix ...string) error {
		f, err := os.CreateTemp(os.TempDir(), "installer-*"+ext)

		if err != nil {
			return fmt.Errorf("unable to create temporary file: %w", err)
		}

		defer func() {
			f.Close()
			os.Remove(f.Name())
		}()

		if _, err := f.Write([]byte(body)); err != nil {
			return fmt.Errorf("error writing to file %s: %w", f.Name(), err)
		}

		args := append(argsPrefix, f.Name())

		cmd := exec.Command(exe, args...)

		output, err := cmd.CombinedOutput()

		if err != nil {
			return fmt.Errorf("error executing command %s: %w", cmd.String(), err)
		}

		fmt.Println(string(output))

		return nil
	}

	switch inst.Executable {
	case "powershell":
		return runScript("powershell.exe", ".ps1", inst.InstallScript, "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-File")
	case "bash":
		// return runScript("bash", ".sh", inst.InstallScript)
	default:
		return fmt.Errorf("unsupported executable: %s", inst.Executable)
	}
	return nil
}
