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

func TestValidateImage(t *testing.T) {
	testCases := []struct {
		desc        string
		fileContent string
		expected    error
	}{
		{
			desc: "invalid image config",
			fileContent: `platform: "windows"
image:
  name: "example image"
  display_name: "example image"
  description: "example image"
  enable_dynamic_app_catalog: true
  use_latest_agent_version: false
  tags:
    - team infra
    - envdev
    - no
  dry_run: false`,
			expected: validator.ErrInvalidTagsCreateImage,
		},
		{
			desc: "invalid image config",
			fileContent: `platform: "windows"
image:
  display_name: "example image"
  description: "example image"
  enable_dynamic_app_catalog: true
  use_latest_agent_version: false
  tags:
    - team:infra
    - env:dev
  dry_run: false`,
			expected: validator.ErrInvalidParametersCreateImage,
		},
		{
			desc: "valid image config",
			fileContent: `platform: "windows"
image:
  name: "example_image"
  display_name: "example image"
  description: "example image"
  enable_dynamic_app_catalog: true
  use_latest_agent_version: false
  tags:
    - team:infra
    - env:dev
  dry_run: false`,
			expected: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.TODO()

			file, err := os.CreateTemp("../../testdata", fmt.Sprintf("test_%s_", tC.desc))

			if err != nil {
				t.Errorf("unable to create temp file: %v", err)
			}

			file.WriteString(tC.fileContent)

			lb := backend.LocalBackend{
				Location: file.Name(),
			}

			configData, err := lb.GetConfig(ctx)

			if err != nil {
				t.Errorf("unable to fetch config data: %v", err)
			}

			err = validator.ValidateImage(ctx, configData)

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
