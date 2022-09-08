package infrastructure

import (
	sharedEvent "ddd/internal/shared/event"
	shared "ddd/internal/shared/infrastructure"
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/event"
	"log"
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
	bus.Subscribe(event.NewNewEvent(domain.NewShoppingCart()), repo)
	return repo
}

func (writer *shoppingCartWriteRepoImpl) New() domain.ShoppingCart {
	return domain.NewShoppingCart()
}

func (writer *shoppingCartWriteRepoImpl) Update(domain.ShoppingCart) (domain.ShoppingCart, error) {
	return domain.NewShoppingCart(), nil
}

func (writer *shoppingCartWriteRepoImpl) Handle(e sharedEvent.DomainEvent) error {
	log.Println("receive event ", e)
	switch e.GetType() {
	}
	return nil
}
