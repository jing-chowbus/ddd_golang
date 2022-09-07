package open_save

import (
	shared "ddd/internal/shared/event"
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/event"
)

type OpenAndSave interface {
	Open(OpenCommand) domain.ShoppingCart
	Save(SaveCommand) (domain.ShoppingCart, error)
}

type openAndSaveImpl struct {
	shoppingCarts domain.ShoppingCarts
	eventBus      shared.EventBus
}

func (openAndSave *openAndSaveImpl) Open(command OpenCommand) domain.ShoppingCart {
	shoppingCart := openAndSave.shoppingCarts.New()
	openAndSave.eventBus.Notify(event.NewNewEvent(shoppingCart))
	return shoppingCart
}

func (openAndSave *openAndSaveImpl) Save(command SaveCommand) (domain.ShoppingCart, error) {
	shoppingCart := command.ShoppingCart
	saved, err := openAndSave.shoppingCarts.FindByID(shoppingCart.ID)
	if err != nil {
		return shoppingCart, err
	}
	saved, err = openAndSave.shoppingCarts.Update(shoppingCart.Merge(saved))
	if err != nil {
		openAndSave.eventBus.Notify(event.NewSaveFailedEvent(saved.ID, err))
		return saved, err
	}
	openAndSave.eventBus.Notify(event.NewSaveSuccessEvent(saved.ID))
	return saved, err
}
