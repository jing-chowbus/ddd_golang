package shared

import (
	"ddd/internal/shared/event"
	"ddd/internal/shared/infrastructure"

	"github.com/google/wire"
)

var InfraSet = wire.NewSet(infrastructure.NewGormDatabase)

var EventBusSet = wire.NewSet(event.NewEventBus)
