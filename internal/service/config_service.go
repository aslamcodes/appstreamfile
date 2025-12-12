package service

import (
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

func ImplementConfig(c *config.Config) error {
	// if err := c.SessionScripts.UpdateSessionScriptConfig(SessionScriptLocation()); err != nil {
	// 	return fmt.Errorf("error configuring session scripts: %w", err)
	// }

	services := &services{
		CatalogSvc: &UpdateStackCatalogSvc{
			Exec: &execx.ExecCommander{},
		},
		FileDeploySvc:     &FileDeploySvc{},
		ImageBuildService: &ImageBuildSvc{},
		InstallerService:  &InstallerSvc{},
	}

	for _, i := range c.Installers {
		fmt.Println("Executing installer with", i.Executable)
		err := services.InstallerService.InstallScript(&i)

		if err != nil {
			return fmt.Errorf("error installing %s: %w", i.Executable+i.InstallScript, err)
		}
	}

	for _, f := range c.Files {
		fmt.Println("Deploying file", f.Path)
		err := services.FileDeploySvc.DeployFile(&f)

		if err != nil {
			return fmt.Errorf("error deploying file %s: %w", f.Path, err)
		}
	}

	for _, catalog := range c.Catalogs {
		if err := services.CatalogSvc.UpdateStackCatalog(catalog); err != nil {
			return fmt.Errorf("error updating stack catalog: %w", err)
		}
	}

	if err := services.ImageBuildService.BuildImage(c.Image); err != nil {
		return fmt.Errorf("error building out image: %w", err)
	}

	return nil
}
