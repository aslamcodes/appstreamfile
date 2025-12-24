package backend_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/backend"
	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockS3Client struct {
	Output *s3.GetObjectOutput
	Err    error
}

func (client *mockS3Client) GetObject(ctx context.Context, bucket string, key string, versionId string) (*s3.GetObjectOutput, error) {
	return client.Output, client.Err
}

func TestS3GetConfig(t *testing.T) {
	client := &mockS3Client{}

	content := `platform: "unix"
installers:
  - executable: "bash"
    installScript: |
      echo "Hello World"`

	expected := &config.Config{
		Platform: "unix",
		Installers: []config.Installer{
			{
				InstallScript: `echo "Hello World"`,
				Executable:    "bash",
			},
		},
	}

	client.Output = &s3.GetObjectOutput{
		Body: io.NopCloser(bytes.NewReader([]byte(content))),
	}

	backend := &backend.S3Backend{
		Bucket:    "test",
		Key:       "test",
		VersionId: "",
		Client:    client,
	}

	actual, err := backend.GetConfig(context.TODO())

	if err != nil {
		t.Errorf("error fetching the config: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v+,\n actual %v+", expected, actual)
	}

}

func TestGetConfigFail(t *testing.T) {
	expectedErr := errors.New("boom")

	backend := &backend.S3Backend{
		Bucket:    "test",
		Key:       "test",
		VersionId: "",
		Client: &mockS3Client{
			Err: expectedErr,
		},
	}

	_, err := backend.GetConfig(context.TODO())

	if err == nil {
		t.Errorf("expected %v, got nil", expectedErr)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected %v, got %v", expectedErr, err)
	}

}
