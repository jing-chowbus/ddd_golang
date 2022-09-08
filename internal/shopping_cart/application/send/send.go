package send

import (
	shared "ddd/internal/shared/entity"
	sharedEvent "ddd/internal/shared/event"
	"ddd/internal/shopping_cart/domain"
	"log"
)

type SendingShoppingCart interface {
	Send(SendCommand) ([]shared.LineItem, error)
}

type sendingShoppingCartImpl struct {
	shoppingCarts domain.ShoppingCarts
	eventBus      sharedEvent.EventBus
}

func NewSendingShoppingCart(shoppingCarts domain.ShoppingCarts, eventBus sharedEvent.EventBus) SendingShoppingCart {
	return &sendingShoppingCartImpl{
		shoppingCarts,
		eventBus,
	}
}

func (checkout *sendingShoppingCartImpl) Send(command SendCommand) ([]shared.LineItem, error) {
	log.Panicln("not implemented")
	return []shared.LineItem{}, nil
}

func NewCheckoutShoppingCart() SendingShoppingCart {
	return &sendingShoppingCartImpl{}
}
