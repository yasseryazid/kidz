package talent

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/yasseryazid/boilerpart/internal/domain/talent"
	presenter "github.com/yasseryazid/boilerpart/internal/presenter/talent"
	usecase "github.com/yasseryazid/boilerpart/internal/usecase/talent"
)

type createTalentRequest struct {
	Name   string   `json:"name" binding:"required"`
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
}

const (
	errListTalents    = "failed to list talents"
	errInvalidRequest = "invalid request"
	errInvalidTalent  = "invalid talent"
	errCreateTalent   = "failed to create talent"
	errInvalidID      = "invalid talent id"
	errTalentNotFound = "talent not found"
	errFetchTalent    = "failed to fetch talent"
)

// RegisterRoutes wires talent endpoints to the router.
func RegisterRoutes(r *gin.Engine, svc *usecase.Service) {
	r.GET("/talents", listTalents(svc))
	r.POST("/talents", createTalent(svc))
	r.GET("/talents/:id", getTalent(svc))
}

func listTalents(svc *usecase.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		talents, err := svc.List(c.Request.Context())
		if err != nil {
			presenter.RespondError(c, http.StatusInternalServerError, errListTalents)
			return
		}

		presenter.RespondData(c, http.StatusOK, talents)
	}
}

func createTalent(svc *usecase.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createTalentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			presenter.RespondError(c, http.StatusBadRequest, errInvalidRequest)
			return
		}

		talent, err := svc.Create(c.Request.Context(), req.Name, req.Title, req.Skills)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidTalent) {
				presenter.RespondError(c, http.StatusBadRequest, errInvalidTalent)
				return
			}

			presenter.RespondError(c, http.StatusInternalServerError, errCreateTalent)
			return
		}

		presenter.RespondData(c, http.StatusCreated, talent)
	}
}

func getTalent(svc *usecase.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		talent, err := svc.GetByID(c.Request.Context(), id)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrInvalidID):
				presenter.RespondError(c, http.StatusBadRequest, errInvalidID)
			case errors.Is(err, domain.ErrNotFound):
				presenter.RespondError(c, http.StatusNotFound, errTalentNotFound)
			default:
				presenter.RespondError(c, http.StatusInternalServerError, errFetchTalent)
			}
			return
		}

		presenter.RespondData(c, http.StatusOK, talent)
	}
}
