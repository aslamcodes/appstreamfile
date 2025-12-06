package config

type Installer struct {
	InstallScript string `yaml:"installScript"`
	Executable    string `yaml:"executable"`
}

func (i *Installer) Install() error {
	return PlatformInstaller(i)
}
