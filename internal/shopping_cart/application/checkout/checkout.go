package checkout

import "ddd/internal/shopping_cart/domain"

type CheckoutShoppingCart interface {
	Checkout(CheckoutCommand) (domain.ShoppingCart, error)
}
