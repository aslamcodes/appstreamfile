package service

import (
	"fmt"

	"github.com/aslamcodes/appstreamfile/internal/config"
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

func Setup(c *config.Config) error {
	// if err := c.SessionScripts.UpdateSessionScriptConfig(SessionScriptLocation()); err != nil {
	// 	return fmt.Errorf("error configuring session scripts: %w", err)
	// }

	// for _, f := range c.Files {
	// 	fmt.Println("Deploying file", f.Path)
	// 	err := f.Deploy()

	// 	if err != nil {
	// 		return fmt.Errorf("error deploying file %s: %w", f.Path, err)
	// 	}
	// }

	// for _, i := range c.Installers {
	// 	fmt.Println("Executing installer with", i.Executable)
	// 	err := i.Install()

	// 	if err != nil {
	// 		return fmt.Errorf("error installing %s: %w", i.Executable+i.InstallScript, err)
	// 	}
	// }
	//
	services := &services{
		CatalogSvc: &UpdateStackCatalogSvc{
			Exec: &execx.ExecCommander{},
		},
	}

	for _, catalog := range c.Catalogs {
		if err := services.CatalogSvc.UpdateStackCatalog(catalog); err != nil {
			return fmt.Errorf("error updating stack catalog: %w", err)
		}
	}

	// if err := c.Image.BuildImage(); err != nil {
	// 	return fmt.Errorf("error building out image: %w", err)
	// }

	return nil
}
