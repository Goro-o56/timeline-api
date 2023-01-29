package registry

import (
	"github.com/gin-gonic/gin"
	"timeline-api/app/config"
)

func NewRouter(config config.App) *gin.Engine {

	sh := NewTimelineHandler()
	// Initialize application
	r := gin.Default()
	r.GET("/", sh.Get)
	return r
}
