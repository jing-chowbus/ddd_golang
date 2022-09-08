package event

import eBus "github.com/asaskevich/EventBus"

type Handler interface {
	Handle(DomainEvent) error
}

type EventBus interface {
	Publish(DomainEvent)
	Subscribe(DomainEvent, Handler) error
	SubscribeAsync(DomainEvent, Handler) error
	SubscribeOnce(DomainEvent, Handler) error
	Unsubscribe(DomainEvent, Handler) error
}

type eventBusImpl struct {
	inner eBus.Bus
}

func NewEventBus() EventBus {
	return &eventBusImpl{
		inner: eBus.New(),
	}
}

func (bus *eventBusImpl) Publish(event DomainEvent) {
	bus.inner.Publish(string(event.GetType()), event)
}

func (bus *eventBusImpl) Subscribe(event DomainEvent, hander Handler) error {
	return bus.inner.Subscribe(string(event.GetType()), hander.Handle)
}

func (bus *eventBusImpl) SubscribeAsync(event DomainEvent, hander Handler) error {
	return bus.inner.SubscribeAsync(string(event.GetType()), hander.Handle, false)
}

func (bus *eventBusImpl) SubscribeOnce(event DomainEvent, hander Handler) error {
	return bus.inner.SubscribeOnce(string(event.GetType()), hander.Handle)
}

func (bus *eventBusImpl) Unsubscribe(event DomainEvent, handler Handler) error {
	return bus.inner.Unsubscribe(string(event.GetType()), handler.Handle)
}
