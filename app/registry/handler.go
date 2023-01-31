package registry

import (
	"timeline-api/app/domain/repository"
	"timeline-api/app/interface/api/server"
	"timeline-api/app/presentation/handler"
)

func NewPersonaHandler(r repository.Persona) handler.Persona {
	return server.NewPersonaHandler(r)
}
