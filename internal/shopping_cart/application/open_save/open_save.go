package open_save

import (
	"ddd/internal/shopping_cart/domain"
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
	return saved, err
}
