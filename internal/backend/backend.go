package backend

import c "github.com/aslamcodes/powerappstream-builder/internal/config"

type BackendSource interface {
	GetConfig() c.Config
}
