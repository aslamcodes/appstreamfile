package config

type Config struct {
	Platform       string          `yaml:"platform"`
	Installers     []Installer     `yaml:"installers"`
	Files          []File          `yaml:"files"`
	Catalogs       []CatalogConfig `yaml:"catalog"`
	SessionScripts SessionScripts  `yaml:"session_scripts"`
	Image          Image           `yaml:"image"`
}
