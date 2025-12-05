package backend_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/aslamcodes/powerappstream-builder/internal/backend"
	"github.com/aslamcodes/powerappstream-builder/internal/config"
)

func TestGetConfig(t *testing.T) {
	localBackend := backend.LocalBackend{
		Location: "../../testdata/config_win.yaml",
	}

	actual, err := localBackend.GetConfig()

	if err != nil {
		t.Fatal(err)
		os.Exit(1)
	}

	expected := config.Config{
		Installers: []config.Installer{
			{Executable: "powershell", InstallScript: "Write-Host \"Hello World\"\n"},
			{Executable: "powershell", InstallScript: "echo \"Setting up environment\"\n"},
			{Executable: "cmd.exe", InstallScript: "echo hello world\n"},
		},
		Files: []config.File{
			{Path: `C:\AppStream\Scripts\Start-System.ps1`, Content: "Write-EventLog -LogName Application -Source AppStream -EventID 100 -Message \"System session start\"\n"},
			{Path: `C:\AppStream\Scripts\Start-User.ps1`, Content: "Write-Host \"User profile initialization\"\n"},
			{Path: `C:\AppStream\Scripts\End-System.ps1`, Content: "Write-EventLog -LogName Application -Source AppStream -EventID 200 -Message \"System cleanup\"\n"},
			{Path: `C:\AppStream\Scripts\End-User.ps1`, Content: "Write-Host \"User session cleanup\"\n"},
		},
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("GetConfig() mismatch.\nexpected: %#v\nactual:   %#v", expected, *actual)
		fmt.Fprintln(os.Stderr, "The expected and actual config are not equal")
		t.FailNow()
	}

}
