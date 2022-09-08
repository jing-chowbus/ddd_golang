package domain

import (
	shared "ddd/internal/shared/event"
	"ddd/internal/shared/value"
	"ddd/internal/shopping_cart/event"

	"github.com/google/uuid"
)

type ShoppingCartsQuery struct {
	TimeRange value.TimeRange
}

type ShoppingCartsFinder interface {
	FindByID(uuid.UUID) (ShoppingCart, error)
	FindByQuery(ShoppingCartsQuery) ([]ShoppingCart, int64, error)
}

type ShoppingCartUpdater interface {
	New() ShoppingCart
	Update(ShoppingCart) (ShoppingCart, error)
}

type ShoppingCarts interface {
	ShoppingCartsFinder
	ShoppingCartUpdater
	Publish(shared.DomainEvent)
}

type shoppingCartsImpl struct {
	finder   ShoppingCartsFinder
	updater  ShoppingCartUpdater
	eventBus shared.EventBus
}

func NewShoppingCarts(finder ShoppingCartsFinder, updater ShoppingCartUpdater, eventBus shared.EventBus) ShoppingCarts {
	shoppingCarts := &shoppingCartsImpl{
		finder:   finder,
		updater:  updater,
		eventBus: eventBus,
	}
	return shoppingCarts
}

func (shoppingCarts *shoppingCartsImpl) FindByID(id uuid.UUID) (ShoppingCart, error) {
	return shoppingCarts.finder.FindByID(id)
}

func (shoppingCarts *shoppingCartsImpl) FindByQuery(query ShoppingCartsQuery) ([]ShoppingCart, int64, error) {
	return shoppingCarts.finder.FindByQuery(query)
}

func (shoppingCarts *shoppingCartsImpl) Update(shoppingCart ShoppingCart) (ShoppingCart, error) {
	saved, err := shoppingCarts.updater.Update(shoppingCart)
	return saved, err
}

func (shoppingCarts *shoppingCartsImpl) New() ShoppingCart {
	shoppingCart := NewShoppingCart()
	shoppingCarts.Publish(event.NewNewEvent(shoppingCart))
	return shoppingCart
}

func (shoppingCarts *shoppingCartsImpl) Publish(event shared.DomainEvent, error) {
	shoppingCarts.eventBus.Publish(event)
}
