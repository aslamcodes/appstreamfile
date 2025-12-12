package service

import (
	"fmt"
	"os"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

type InstallerSvc struct {
	Exec           execx.Commander
	KeepTmpForTest bool
}

func (s *InstallerSvc) Install(inst *config.Installer) error {
	var (
		exe  string
		ext  string
		args []string
	)

	switch inst.Executable {
	case "powershell":
		exe = "powershell.exe"
		ext = ".ps1"
		args = []string{"-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-File"}
	case "bash":
		exe = "bash"
		ext = ".sh"
	default:
		return fmt.Errorf("unsupported executable: %s", inst.Executable)
	}

	f, err := os.CreateTemp("", "installer-*"+ext)
	if err != nil {
		return fmt.Errorf("unable to create temporary file: %w", err)
	}
	defer func() {
		if !s.KeepTmpForTest {
			os.Remove(f.Name())
		}
	}()
	defer f.Close()

	_, err = f.Write([]byte(inst.InstallScript))

	if err != nil {
		return fmt.Errorf("writing script: %w", err)
	}

	return s.RunScript(exe, args, f.Name())
}

func (s *InstallerSvc) RunScript(exe string, args []string, filePath string) error {
	cmd := s.Exec.Command(exe, append(args, filePath)...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	fmt.Println(string(out))
	return nil
}
