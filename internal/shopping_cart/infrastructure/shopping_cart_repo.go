package infrastructure

import (
	shared "ddd/internal/shared/infrastructure"
	"ddd/internal/shopping_cart/domain"

	"github.com/google/uuid"
)

type ShoppingCartReadRepo interface {
	domain.ShoppingCartsFinder
}

type shoppingCartReadRepoImpl struct {
	database shared.Database
}

func NewShoppingCartReadRepo(database shared.Database) ShoppingCartReadRepo {
	return &shoppingCartReadRepoImpl{
		database: database,
	}
}

func (read *shoppingCartReadRepoImpl) FindByID(shoppingCartID uuid.UUID) (domain.ShoppingCart, error) {
	return domain.NewShoppingCart(), nil
}

func (read *shoppingCartReadRepoImpl) FindByQuery(domain.ShoppingCartsQuery) ([]domain.ShoppingCart, int64, error) {
	return []domain.ShoppingCart{}, 0, nil
}
