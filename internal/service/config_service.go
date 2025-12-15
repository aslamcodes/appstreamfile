package service

import (
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

const IMAGE_ASSISTANT = "image-assistant"

func ImplementConfig(c *config.Config) error {
	services := &services{
		FileDeploySvc:        &FileDeploySvc{},
		SessionScriptService: &SessionScriptSvc{},
		CatalogSvc: &UpdateStackCatalogSvc{
			Exec: &execx.ExecCommander{},
		},
		ImageBuildService: &ImageBuildSvc{
			Exec: &execx.ExecCommander{},
		},
		InstallerService: &InstallerSvc{
			Exec: &execx.ExecCommander{},
		},
	}

	for _, i := range c.Installers {
		fmt.Println("Executing installer with", i.Executable)
		err := services.InstallerService.InstallScript(&i)

		if err != nil {
			return fmt.Errorf("error executing %s script\n%s: %w", i.Executable, i.InstallScript, err)
		}
	}

	for _, f := range c.Files {
		fmt.Println("Deploying file", f.Path)
		err := services.FileDeploySvc.DeployFile(&f)

		if err != nil {
			return fmt.Errorf("error deploying file %s: %w", f.Path, err)
		}
	}

	if err := services.SessionScriptService.UpdateSessionScriptConfig(SessionScriptLocation(), c.SessionScripts); err != nil {
		return fmt.Errorf("error configuring session scripts: %w", err)
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
