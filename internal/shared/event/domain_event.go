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

func Is(event DomainEvent, target DomainEvent) bool {
	if target == nil || event == nil {
		return false
	}
	return event.GetType() == target.GetType()
}
