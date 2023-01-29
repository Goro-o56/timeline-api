package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"timeline-api/app/handler"
)

type timelineHandler struct{}

func NewTimelineHandler() handler.Timeline {
	return &timelineHandler{}
}

// Get is get application state
// @Summary Return application Timeline
// @Tags Timeline
// @Produce json
// @Success 200 {object} Timeline.State
// @Failure 404 {object} Timeline.Error
// @Failure 405 {object} Timeline.Error
// @Router /v1 [get
func (h *timelineHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Timeline{
		Status:      "Active",
		Environment: gin.Mode(),
		LogLevel:    logrus.GetLevel().String(),
		TimeZone:    time.Local.String(),
	})
}

// NoRoute is not found response
func (h *timelineHandler) NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, entity.Error{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	})
}

// NoMethod is method not allowed response
func (h *timelineHandler) NoMethod(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, entity.Error{
		Code:    http.StatusMethodNotAllowed,
		Message: http.StatusText(http.StatusMethodNotAllowed),
	})
}
