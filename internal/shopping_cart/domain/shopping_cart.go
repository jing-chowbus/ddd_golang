package domain

import (
	sharedEntity "ddd/internal/shared/entity"
	"ddd/internal/shared/value"
	"ddd/internal/shopping_cart/entity"

	"github.com/google/uuid"
)

type ShoppingCart struct {
	ID               uuid.UUID
	Items            []sharedEntity.LineItem
	PretaxDiscounts  []sharedEntity.Discount
	PretaxFees       []sharedEntity.Fee
	PosttaxDiscounts []sharedEntity.Discount
	PosttaxFees      []sharedEntity.Fee
	Payments         []sharedEntity.Payment
	Guests           []entity.Guest
}

func (shoppingCart ShoppingCart) applyOrderLevelDiscountsAndFees() ShoppingCart {
	return shoppingCart
}

func (shoppingCart ShoppingCart) GetSubtotal() value.Money {
	shoppingCart = shoppingCart.applyOrderLevelDiscountsAndFees()
	subtotal := value.Money{}
	for _, item := range shoppingCart.Items {
		subtotal = item.GetSubtotal()
	}
	return subtotal
}

func (shoppingCart ShoppingCart) GetTax() value.Money {
	shoppingCart = shoppingCart.applyOrderLevelDiscountsAndFees()
	tax := value.Money{}
	for _, item := range shoppingCart.Items {
		tax = item.GetTax()
	}
	return tax
}

func (shoppingCart ShoppingCart) GetTotal() value.Money {
	shoppingCart = shoppingCart.applyOrderLevelDiscountsAndFees()
	total := value.Money{}
	for _, item := range shoppingCart.Items {
		total = total.Add(item.GetTotal())
	}
	return total
}

func (shoppingCart ShoppingCart) Merge(other ShoppingCart) ShoppingCart {
	return shoppingCart
}

func NewShoppingCart() ShoppingCart {
	return ShoppingCart{
		ID: uuid.New(),
	}
}
