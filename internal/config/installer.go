package config

type Installer struct {
	InstallScript string `yaml:"installScript"`
	Executable    string `yaml:"executable"`
}
