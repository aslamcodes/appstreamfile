package validator_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/backend"
	"github.com/aslamcodes/appstreamfile/internal/validator"
)

func TestInstallerValidator(t *testing.T) {
	testCases := []struct {
		desc        string
		fileContent string
		expected    error
	}{
		{
			desc: "Invalid Platform",
			fileContent: `platform: "non_existent"
installers:
  - executable: "bash"
    installScript: |
      echo "Hello World"`,
			expected: validator.ErrInvalidPlatform,
		},
		{
			desc: "Invalid Executable",
			fileContent: `platform: "windows"
installers:
  - executable: "bash"
    installScript: |
      echo "Hello World"`,
			expected: validator.ErrInvalidExecutableForPlatform,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			file, err := os.CreateTemp("../../testdata", fmt.Sprintf("test_%s_", tC.desc))

			if err != nil {
				t.Errorf("unable to create temp file: %v", err)
			}

			file.WriteString(tC.fileContent)

			lb := backend.LocalBackend{
				Location: file.Name(),
			}
ctx := context.TODO()
			configData, err := lb.GetConfig(ctx)

			if err != nil {
				t.Errorf("unable to fetch config data: %v", err)
			}

			err = validator.InstallerValidator(ctx,configData)

			if !errors.Is(err, tC.expected) {
				t.Errorf("expected %v, got %v", tC.expected, err)
			}

			t.Cleanup(func() {
				os.Remove(file.Name())
			})

			t.Cleanup(func() {
				file.Close()
			})
		})
	}
}
