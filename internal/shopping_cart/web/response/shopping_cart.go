package response

import (
	sharedEntity "ddd/internal/shared/entity"
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/entity"

	"github.com/google/uuid"
)

type ShoppingCartResponse struct {
	ID uuid.UUID
	// TODO create response objects
	Items            []sharedEntity.LineItem
	PretaxDiscounts  []sharedEntity.Discount
	PretaxFees       []sharedEntity.Fee
	PosttaxDiscounts []sharedEntity.Discount
	PosttaxFees      []sharedEntity.Fee
	Payments         []sharedEntity.Payment
	Guests           []entity.Guest
}

func NewShoppingCartResponse(shoppingCart domain.ShoppingCart) ShoppingCartResponse {
	return ShoppingCartResponse{
		ID:    shoppingCart.ID,
		Items: shoppingCart.Items,
	}
}
