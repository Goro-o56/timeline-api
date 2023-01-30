package repository

import (
	"time"
	"timeline-api/app/domain/entity"
)

type Persona interface {
	Exists(account string) (bool, error)
	Find(id uint) (*entity.Persona, error)
	FindByTime(time time.Time) (*entity.Persona, error)
}
