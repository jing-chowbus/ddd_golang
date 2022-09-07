package entity

import (
	"ddd/internal/shared/value"

	"github.com/google/uuid"
)

type Payment struct {
	ID     uuid.UUID
	Amount value.Money
}
