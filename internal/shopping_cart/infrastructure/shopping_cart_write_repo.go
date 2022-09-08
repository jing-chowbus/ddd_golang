package infrastructure

import (
	sharedEvent "ddd/internal/shared/event"
	shared "ddd/internal/shared/infrastructure"
	"ddd/internal/shopping_cart/domain"
)

type ShoppingCartWriteRepo interface {
	domain.ShoppingCartUpdater
	sharedEvent.Handler
}

type shoppingCartWriteRepoImpl struct {
	database shared.Database
	bus      sharedEvent.EventBus
}

func NewShoppingCartWriteRepo(database shared.Database, bus sharedEvent.EventBus) ShoppingCartWriteRepo {
	repo := &shoppingCartWriteRepoImpl{
		database: database,
		bus:      bus,
	}
	return repo
}

func (writer *shoppingCartWriteRepoImpl) New() domain.ShoppingCart {
	return domain.NewShoppingCart()
}

func (writer *shoppingCartWriteRepoImpl) Update(domain.ShoppingCart) (domain.ShoppingCart, error) {
	return domain.NewShoppingCart(), nil
}

func (writer *shoppingCartWriteRepoImpl) Handle(e sharedEvent.DomainEvent) error {
	switch e.GetType() {
	}
	return nil
}
