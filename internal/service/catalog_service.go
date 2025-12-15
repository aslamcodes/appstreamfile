package service

import (
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

type UpdateStackCatalogSvc struct {
	Exec execx.Commander
}

func (svc *UpdateStackCatalogSvc) UpdateStackCatalog(c config.CatalogConfig) error {
	fmt.Println("\nConfiguring stack catalog")
	fmt.Println(c)

	_, err := svc.Exec.LookPath(IMAGE_ASSISTANT)

	if err != nil {
		return err
	}

	args := append([]string{"add-application"}, c.Args()...)

	cmd := svc.Exec.Command(IMAGE_ASSISTANT, args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
