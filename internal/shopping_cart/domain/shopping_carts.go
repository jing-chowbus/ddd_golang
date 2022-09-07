package domain

import (
	"ddd/internal/shared/event"
	shared "ddd/internal/shared/event"
	"ddd/internal/shared/value"

	"github.com/google/uuid"
)

type ShoppingCartsQuery struct {
	TimeRange value.TimeRange
}

type ShoppingCartsFinder interface {
	FindByID(uuid.UUID) (ShoppingCart, error)
	FindByQuery(ShoppingCartsQuery) ([]ShoppingCart, error)
}

type ShoppingCartUpdater interface {
	New() ShoppingCart
	Update(ShoppingCart) (ShoppingCart, error)
}

type ShoppingCarts interface {
	ShoppingCartsFinder
	ShoppingCartUpdater
	Notify(shared.DomainEvent) error
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

func (shoppingCarts *shoppingCartsImpl) FindByQuery(query ShoppingCartsQuery) ([]ShoppingCart, error) {
	return shoppingCarts.finder.FindByQuery(query)
}

func (shoppingCarts *shoppingCartsImpl) Update(shoppingCart ShoppingCart) (ShoppingCart, error) {
	saved, err := shoppingCarts.updater.Update(shoppingCart)
	return saved, err
}

func (shoppingCarts *shoppingCartsImpl) New() ShoppingCart {
	return NewShoppingCart()
}

func (shoppingCarts *shoppingCartsImpl) Notify(event event.DomainEvent) error {
	return shoppingCarts.eventBus.Notify(event)
}
