package event

import (
	"ddd/internal/shared/exceptoins"
	"sync"
)

type Subscriber interface {
	Subscribe(EventType, Handler) error
}

type Notifier interface {
	Notify(DomainEvent) error
	NotifyAll([]DomainEvent) error
}

type Handler interface {
	handle(DomainEvent) error
}

type EventBus interface {
	Notifier
	Subscriber
}

type eventBusImpl struct {
	mu          sync.Mutex
	subscribers map[EventType][]Handler
}

func (bus *eventBusImpl) Notify(event DomainEvent) error {
	handlers := bus.subscribers[event.GetType()]
	for _, handler := range handlers {
		if err := handler.handle(event); err != nil {
			return err
		}
	}
	return nil
}

func (bus *eventBusImpl) NotifyAll(events []DomainEvent) error {
	for _, event := range events {
		if err := bus.Notify(event); err != nil {
			return err
		}
	}
	return nil
}

func (bus *eventBusImpl) Subscribe(event DomainEvent, handler Handler) error {
	locked := bus.mu.TryLock()
	if !locked {
		return exceptoins.SystemException{}
	}
	defer bus.mu.Unlock()
	if handlers, ok := bus.subscribers[event.GetType()]; !ok {
		bus.subscribers[event.GetType()] = []Handler{handler}
	} else {
		bus.subscribers[event.GetType()] = append(handlers, handler)
	}
	return nil
}
