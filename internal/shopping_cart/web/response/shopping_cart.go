package response

import (
	"ddd/internal/shopping_cart/domain"

	"github.com/google/uuid"
)

type ShoppingCartResponse struct {
	ID uuid.UUID
}

func NewShoppingCartResponse(shoppingCart domain.ShoppingCart) ShoppingCartResponse {
	return ShoppingCartResponse{
		ID: shoppingCart.ID,
	}
}
