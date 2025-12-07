package validator_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aslamcodes/appstream-file/internal/backend"
	"github.com/aslamcodes/appstream-file/internal/validator"
)

func TestFileDeployValidator(t *testing.T) {
	testCases := []struct {
		desc        string
		fileContent string
		expected    error
	}{
		{
			desc: "invalid file deploy without path",
			fileContent: `platform: "windows"
files:
- path: "C:/Appstream/session"
  content: |
    echo "Hello World"

- content: |
    echo "Hello World"`,
			expected: validator.ErrFileDeployPathMissing,
		},
		{
			desc: "valid file deploy configuration",
			fileContent: `platform: "unix"
files:
  - path: "C:/Appstream/session"
    content: |
      echo "Hello World"`,
			expected: nil,
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

			err = validator.ValidateFileDeploys(configData)

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
