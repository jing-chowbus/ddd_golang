package infrastructure

import (
	shared "ddd/internal/shared/infrastructure"
	"ddd/internal/shopping_cart/domain"
)

type ShoppingCartReadRepo interface {
	domain.ShoppingCartsFinder
}

type ShoppingCartWriteRepo interface {
	domain.ShoppingCartUpdater
}

type shoppingCartReadRepoImpl struct {
	database shared.Database
}

type shoppingCartWriteRepoImpl struct {
	database shared.Database
}

func NewShoppingCartReadRepo(database shared.Database) ShoppingCartReadRepo {
	return &shoppingCartReadRepoImpl{
		database: database,
	}
}

func NewShoppingCartWriteRepo(database shared.Database) ShoppingCartReadRepo {
	return &shoppingCartReadRepoImpl{
		database: database,
	}
}
