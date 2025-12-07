package validator_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aslamcodes/powerappstream-builder/internal/backend"
	"github.com/aslamcodes/powerappstream-builder/internal/validator"
)

func TestValidatePaths(t *testing.T) {
	testCases := []struct {
		desc        string
		fileContent string
		expected    error
	}{
		{
			desc: "invalid config windows",
			fileContent: `platform: "windows"
files:
  - path: "/var/log"
    content: |
      echo "Hello World"`,
			expected: validator.ErrInvalidPathForWindows,
		},
		{
			desc: "invalid config unix",
			fileContent: `platform: "unix"
files:
  - path: "C:/Appstream/session"
    content: |
      echo "Hello World"`,
			expected: validator.ErrInvalidPathForUnix,
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

			configData, err := lb.GetConfig()

			if err != nil {
				t.Errorf("unable to fetch config data: %v", err)
			}

			err = validator.ValidatePaths(configData)

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
