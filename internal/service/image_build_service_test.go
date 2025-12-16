package service_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/service"
)

func TestImageBuild(t *testing.T) {
	fc := FakeCommander{
		LastCommand: "",
		LastArgs:    []string{},
		LookPathErr: nil,
	}

	i := service.ImageBuildSvc{
		Exec: &fc,
	}

	image := config.Image{
		Name:                    "test",
		DisplayName:             "test2",
		Description:             "test3",
		EnableDynamicAppCatalog: true,
		UseLatestAgentVersion:   true,
		Tags:                    []string{"k1", "v1", "k2", "built with appstreamfile"},
		DryRun:                  true,
	}

	i.BuildImage(image)

	expectedCommand := strings.TrimSpace(`image-assistant create-image --name test --display-name test2 --description test3 --use-latest-agent-version --enable-dynamic-app-catalog --dry-run --tags k1 v1 k2 built with appstreamfile`)
	actual := strings.TrimSpace(fmt.Sprintf("%s %s", fc.LastCommand, strings.Join(fc.LastArgs, " ")))

	if expectedCommand != actual {
		t.Errorf("\nexpected %s,\ngot %s", expectedCommand, actual)
	}

}
