package health

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Registrar registers health-related routes.
type Registrar struct {
	db *sql.DB
}

func NewRegistrar(db *sql.DB) *Registrar {
	return &Registrar{db: db}
}

func (r *Registrar) Register(router *gin.Engine) {
	router.GET("/ping", Ping)
	if r.db != nil {
		router.GET("/db-ping", DBPing(r.db))
	}
}
