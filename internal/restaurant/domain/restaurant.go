package domain

import (
	"ddd/internal/restaurant/entity"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID         uuid.UUID
	Info       entity.Info
	Preference entity.Preference
}
