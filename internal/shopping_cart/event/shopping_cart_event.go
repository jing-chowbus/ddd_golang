package event

import (
	"ddd/internal/shared/event"
	"time"

	"github.com/google/uuid"
)

type ShoppingCartEvent interface {
	event.DomainEvent
}

type shoppingCartEventImpl struct {
	eventType      event.EventType
	shoppingCartID uuid.UUID
	when           time.Time
}

func (impl shoppingCartEventImpl) GetAggrateID() uuid.UUID {
	return impl.shoppingCartID
}

func (impl shoppingCartEventImpl) GetEventTime() time.Time {
	return impl.when
}

func (impl shoppingCartEventImpl) GetType() event.EventType {
	return impl.eventType
}

type NewEvent struct {
	shoppingCartEventImpl
}

func NewNewEvent(shoppingCartID uuid.UUID) NewEvent {
	return NewEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:new",
			shoppingCartID: shoppingCartID,
		},
	}
}

type SaveEvent struct {
	shoppingCartEventImpl
}

func NewSaveEvent(shoppingCartID uuid.UUID) SaveEvent {
	return SaveEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:save",
			shoppingCartID: shoppingCartID,
		},
	}
}

type SaveSuccessEvent struct {
	shoppingCartEventImpl
}

func NewSaveSuccessEvent(shoppingCartID uuid.UUID) SaveSuccessEvent {
	return SaveSuccessEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:save:success",
			shoppingCartID: shoppingCartID,
		},
	}
}

type SaveFailedEvent struct {
	shoppingCartEventImpl
	err error
}

func NewSaveFailedEvent(shoppingCartID uuid.UUID, err error) SaveFailedEvent {
	return SaveFailedEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:save:failed",
			shoppingCartID: shoppingCartID,
		},
		err: err,
	}
}

type CheckoutEvent struct {
	shoppingCartEventImpl
}

func NewCheckoutEvent(shoppingCartID uuid.UUID) CheckoutEvent {
	return CheckoutEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:checkout",
			shoppingCartID: shoppingCartID,
		},
	}
}

type CheckoutSuccessEvent struct {
	shoppingCartEventImpl
}

func NewCheckoutSuccessEvent(shoppingCartID uuid.UUID) CheckoutSuccessEvent {
	return CheckoutSuccessEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:checkout:success",
			shoppingCartID: shoppingCartID,
		},
	}
}

type CheckoutFailedEvent struct {
	shoppingCartEventImpl
	err error
}

func NewCheckoutFailedEvent(shoppingCartID uuid.UUID) CheckoutFailedEvent {
	return CheckoutFailedEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			eventType:      "shopping_cart:checkout:failed",
			shoppingCartID: shoppingCartID,
		},
	}
}

var _ ShoppingCartEvent = NewEvent{}
var _ ShoppingCartEvent = SaveEvent{}
var _ ShoppingCartEvent = SaveSuccessEvent{}
var _ ShoppingCartEvent = SaveFailedEvent{}
var _ ShoppingCartEvent = CheckoutEvent{}
var _ ShoppingCartEvent = CheckoutSuccessEvent{}
var _ ShoppingCartEvent = CheckoutFailedEvent{}
