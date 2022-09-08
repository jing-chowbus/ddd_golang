package open_save

import (
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/event"
)

type OpenAndSave interface {
	Open(OpenCommand) domain.ShoppingCart
	Save(SaveCommand) (domain.ShoppingCart, error)
}

type openAndSaveImpl struct {
	shoppingCarts domain.ShoppingCarts
}

func NewOpenAndSave(shoppingCarts domain.ShoppingCarts) OpenAndSave {
	return &openAndSaveImpl{
		shoppingCarts,
	}
}

func (openAndSave *openAndSaveImpl) Open(command OpenCommand) domain.ShoppingCart {
	shoppingCart := openAndSave.shoppingCarts.New()
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
		openAndSave.eventBus.Publish(event.NewSaveFailedEvent(saved.ID, err))
		return saved, err
	}
	openAndSave.eventBus.Publish(event.NewSaveSuccessEvent(saved.ID))
	return saved, err
}
