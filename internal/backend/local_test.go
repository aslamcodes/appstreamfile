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
		Location: "../../testdata/config.yaml",
	}

	actual, err := localBackend.GetConfig()

	if err != nil {
		t.Fatal(err)
		os.Exit(1)
	}

	expected := config.Config{
		Installers: []config.Installer{
			{
				Executable:    "powershell",
				InstallScript: "Write-Host \"Hello World\"\n",
			},
			{
				Executable:    "powershell",
				InstallScript: "echo \"Setting up environment\"\napt-get update\napt-get install -y curl\n",
			},
		},
}

	if !reflect.DeepEqual(expected, *actual) {
		fmt.Fprintln(os.Stderr, "The expected and actual config are not equal")
		t.FailNow()
	}

}
