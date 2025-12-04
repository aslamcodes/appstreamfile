package agent

import (
	"io"

	c "github.com/aslamcodes/powerappstream-builder/internal/config"
)

type Agent struct{}

func (a *Agent) HandleConfig(config *c.Config, out io.Writer) {
	for _, installer := range config.Installers {
		installer.Install(out)
	}
}
