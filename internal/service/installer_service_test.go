package service_test

import (
	"os"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/service"
)

func TestInstallCreatesScriptFile(t *testing.T) {
	fake := &FakeCommander{}
	svc := &service.InstallerSvc{Exec: fake, KeepTmpForTest: true}

	inst := &config.Installer{
		Executable:    "powershell",
		InstallScript: "Write-Host hi",
	}

	_ = svc.InstallScript(inst)

	scriptPath := fake.LastArgs[len(fake.LastArgs)-1]

	b, _ := os.ReadFile(scriptPath)

	if string(b) != "Write-Host hi" {
		t.Logf("script path %s", scriptPath)
		t.Errorf("expected %s\ngot %s", inst.InstallScript, string(b))
	}
}

func TestRunScript(t *testing.T) {
	fake := &FakeCommander{}
	svc := &service.InstallerSvc{Exec: fake}

	err := svc.RunScript("powershell.exe",
		[]string{"-NoProfile", "-File"},
		"/tmp/test.ps1",
	)
	if err != nil {
		t.Fatal(err)
	}

	if fake.LastCommand != "powershell.exe" {
		t.Fatalf("wrong program: %s", fake.LastCommand)
	}
	if fake.LastArgs[len(fake.LastArgs)-1] != "/tmp/test.ps1" {
		t.Fatalf("missing script path")
	}
}
