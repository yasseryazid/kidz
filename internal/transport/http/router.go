package http

import "github.com/gin-gonic/gin"

// RouteRegistrar defines a module responsible for registering routes.
type RouteRegistrar interface {
	Register(*gin.Engine)
}

// NewRouter builds the HTTP router with registered modules.
func NewRouter(registrars ...RouteRegistrar) *gin.Engine {
	r := gin.Default()
	for _, registrar := range registrars {
		if registrar == nil {
			continue
		}
		registrar.Register(r)
	}

	return r
}
