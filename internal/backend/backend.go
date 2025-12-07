package backend

import c "github.com/aslamcodes/appstream-file/internal/config"

type BackendSource interface {
	GetConfig() c.Config
}
