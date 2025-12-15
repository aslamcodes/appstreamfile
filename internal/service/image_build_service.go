package service

import (
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

type ImageBuildSvc struct {
	Exec execx.Commander
}

func (i *ImageBuildSvc) BuildImage(image config.Image) error {
	fmt.Println("\nBuilding out image")

	_, err := i.Exec.LookPath(IMAGE_ASSISTANT)

	if err != nil {
		return err
	}

	args := append([]string{"create-image"}, image.Args()...)

	cmd := i.Exec.Command(IMAGE_ASSISTANT, args...)

	fmt.Println(cmd.String())

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
