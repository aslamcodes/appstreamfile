package validator_test

import (
	"testing"

	"github.com/aslamcodes/appstream-file/internal/backend"
	"github.com/aslamcodes/appstream-file/internal/validator"
)

func TestValidateCatalogConfig(t *testing.T) {
	testCases := []struct {
		desc     string
		filename string
		expected error
	}{
		{
			desc:     "Config name and path must be given",
			filename: "../../testdata/win_valid_config.yaml",
			expected: nil,
		},
		{
			desc:     "Invalid Config path must be given",
			filename: "../../testdata/win_invalid_config_path.yaml",
			expected: validator.ErrEmptyCatalogPath,
		},
		{
			desc:     "Invalid Config name and path must be given",
			filename: "../../testdata/win_invalid_config_name.yaml",
			expected: validator.ErrEmptyCatalogName,
		},
		{
			desc:     "Invalid Config name and path must be given",
			filename: "../../testdata/win_invalid_config_name_path.yaml",
			expected: validator.ErrEmptyCatalogNamePath,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lb := backend.LocalBackend{
				Location: tC.filename,
			}

			config, err := lb.GetConfig()

			if err != nil {
				t.Errorf("unable to load config: %s", err.Error())
			}

			actual := validator.ValidateCatalogApplications(config)

			if actual != tC.expected {
				t.Errorf("expected %v, actual %v", tC.expected, actual)
			}
		})
	}
}
