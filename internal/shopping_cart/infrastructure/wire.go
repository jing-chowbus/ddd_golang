package infrastructure

import (
	"ddd/internal/shopping_cart/domain"

	"github.com/google/wire"
)

var InfraSet = wire.NewSet(NewShoppingCartReadRepo, NewShoppingCartWriteRepo, wire.Bind(new(domain.ShoppingCartsFinder), new(ShoppingCartReadRepo)), wire.Bind(new(domain.ShoppingCartUpdater), new(ShoppingCartWriteRepo)))
