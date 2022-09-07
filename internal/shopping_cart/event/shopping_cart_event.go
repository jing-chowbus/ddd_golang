package event

import (
	"ddd/internal/shared/event"
	"ddd/internal/shopping_cart/domain"
	"time"

	"github.com/google/uuid"
)

type ShoppingCartEvent interface {
	event.DomainEvent
}

type shoppingCartEventImpl struct {
	shoppingCartID uuid.UUID
	when           time.Time
}

func (impl shoppingCartEventImpl) GetAggrateID() uuid.UUID {
	return impl.shoppingCartID
}

func (impl shoppingCartEventImpl) GetEventTime() time.Time {
	return impl.when
}

var _ ShoppingCartEvent = NewEvent{}
var _ ShoppingCartEvent = SaveEvent{}
var _ ShoppingCartEvent = SaveSuccessEvent{}
var _ ShoppingCartEvent = SaveFailedEvent{}
var _ ShoppingCartEvent = CheckoutEvent{}
var _ ShoppingCartEvent = CheckoutSuccessEvent{}
var _ ShoppingCartEvent = CheckoutFailedEvent{}

type NewEvent struct {
	shoppingCartEventImpl
	ShoppingCart domain.ShoppingCart
}

func (NewEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.new")
}

func NewNewEvent(shoppingCart domain.ShoppingCart) NewEvent {
	return NewEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			shoppingCartID: shoppingCart.ID,
		},
		ShoppingCart: shoppingCart,
	}
}

type SaveEvent struct {
	shoppingCartEventImpl
	ShoppingCart domain.ShoppingCart
}

func (SaveEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.save")
}

func NewSaveEvent(shoppingCart domain.ShoppingCart) SaveEvent {
	return SaveEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			shoppingCartID: shoppingCart.ID,
		},
		ShoppingCart: shoppingCart,
	}
}

type SaveSuccessEvent struct {
	shoppingCartEventImpl
}

func (SaveSuccessEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.save.success")
}

func NewSaveSuccessEvent(shoppingCartID uuid.UUID) SaveSuccessEvent {
	return SaveSuccessEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			shoppingCartID: shoppingCartID,
		},
	}
}

type SaveFailedEvent struct {
	shoppingCartEventImpl
	err error
}

func (SaveFailedEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.save.failed")
}

func NewSaveFailedEvent(shoppingCartID uuid.UUID, err error) SaveFailedEvent {
	return SaveFailedEvent{
		shoppingCartEventImpl: shoppingCartEventImpl{
			shoppingCartID: shoppingCartID,
		},
		err: err,
	}
}

type CheckoutEvent struct {
	shoppingCartEventImpl
	ShoppingCart domain.ShoppingCart
}

func (CheckoutEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.checkout")
}

type CheckoutSuccessEvent struct {
	shoppingCartEventImpl
}

func (CheckoutSuccessEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.checkout.success")
}

type CheckoutFailedEvent struct {
	shoppingCartEventImpl
	err error
}

func (CheckoutFailedEvent) GetType() event.EventType {
	return event.EventType("shopping_cart.checkout.failed")
}
