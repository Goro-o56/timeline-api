package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"timeline-api/app/domain/entity"
	"timeline-api/app/domain/repository"
	"timeline-api/app/presentation/handler"
)

type personaHandler struct {
	repo repository.Persona
}

func NewPersonaHandler(ur repository.Persona) handler.Persona {
	return &personaHandler{
		repo: ur,
	}
}
func (p *personaHandler) Find(c *gin.Context) {
	//TODO ここの実装、まだ変数を置けていない
	c.JSON(http.StatusOK, entity.Persona{
		ID:        0,
		Content:   "",
		CreatedAt: time.Time{},
	})
}
