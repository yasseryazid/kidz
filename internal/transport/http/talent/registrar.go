package talent

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/yasseryazid/boilerpart/internal/usecase/talent"
)

// Registrar registers talent-related routes.
type Registrar struct {
	service *usecase.Service
}

func NewRegistrar(service *usecase.Service) *Registrar {
	return &Registrar{service: service}
}

func (r *Registrar) Register(router *gin.Engine) {
	if r.service == nil {
		return
	}
	RegisterRoutes(router, r.service)
}
