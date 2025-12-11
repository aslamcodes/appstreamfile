package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/service"
)

func TestCatalogConfigSuccess(t *testing.T) {
	fc := &FakeCommander{
		LookPathErr: nil,
	}

	catalog := config.CatalogConfig{
		Name:        "Notepad",
		Path:        "C://System32//notepad.exe",
		DisplayName: "Notepad",
		Parameters:  "--file hello.txt",
		IconPath:    "C://System32//notepad.ico",
		WorkingDir:  ".",
	}

	updateCatalogSvc := service.UpdateStackCatalogSvc{
		Exec: fc,
	}

	updateCatalogSvc.UpdateStackCatalog(catalog)

	if !slices.Equal(catalog.Args(), fc.LastArgs[1:]) {
		t.Errorf("expected %v\n, got %v", catalog.Args(), fc.LastArgs[1:])
	}

}

func TestCatalogConfigFail(t *testing.T) {
	fc := &FakeCommander{
		LookPathErr: errors.New("file not found"),
	}

	catalog := config.CatalogConfig{
		Name:        "Notepad",
		Path:        "C://System32//notepad.exe",
		DisplayName: "Notepad",
		Parameters:  "--file hello.txt",
		IconPath:    "C://System32//notepad.ico",
		WorkingDir:  ".",
	}

	updateCatalogSvc := service.UpdateStackCatalogSvc{
		Exec: fc,
	}

	err := updateCatalogSvc.UpdateStackCatalog(catalog)

	if !errors.Is(fc.LookPathErr, err) {
		t.Errorf("expected %v\n, got %v", fc.LookPathErr, err)
	}

}
