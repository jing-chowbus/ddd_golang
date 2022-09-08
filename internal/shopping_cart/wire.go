package shoppingcart

import (
	"ddd/internal/shopping_cart/application/checkout"
	"ddd/internal/shopping_cart/application/open_save"
	"ddd/internal/shopping_cart/application/send"
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/infrastructure"
	"ddd/internal/shopping_cart/web"

	"github.com/google/wire"
)

var domainSet = wire.NewSet(infrastructure.InfraSet, domain.NewShoppingCarts)
var applicationSet = wire.NewSet(domainSet, checkout.NewCheckoutShoppingCart, open_save.NewOpenAndSave, send.NewSendingShoppingCart)
var ShoppCartWeb = wire.NewSet(applicationSet, web.NewShopingCartsWebHanders)
