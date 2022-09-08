package checkout

import (
	"ddd/internal/shopping_cart/domain"
	"log"
)

type CheckoutShoppingCart interface {
	Checkout(CheckoutCommand) (domain.ShoppingCart, error)
}

type checkoutShoppingCartImpl struct {
}

func (checkout *checkoutShoppingCartImpl) Checkout(command CheckoutCommand) (domain.ShoppingCart, error) {
	log.Panicln("not implemented")
	return domain.NewShoppingCart(), nil
}

func NewCheckoutShoppingCart() CheckoutShoppingCart {
	return &checkoutShoppingCartImpl{}
}
