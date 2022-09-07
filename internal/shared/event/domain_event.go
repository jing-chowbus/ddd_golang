package event

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

type DomainEvent interface {
	GetAggrateID() uuid.UUID
	GetEventTime() time.Time
	GetType() EventType
}
