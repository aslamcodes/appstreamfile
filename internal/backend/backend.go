package backend

import c "github.com/aslamcodes/appstreamfile/internal/config"

type BackendSource interface {
	GetConfig() c.Config
}
