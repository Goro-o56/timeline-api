package registry

import (
	"timeline-api/app/handler"
	"timeline-api/app/interface/api/server"
)

func NewTimelineHandler() handler.Timeline {
	return server.NewTimelineHandler()
}
