package entity

import (
	"ddd/internal/shared/value"

	"github.com/google/uuid"
)

type Discount struct {
	ID         uuid.UUID
	OrderID    uuid.UUID
	LineItemID uuid.UUID
	BeforeTax  bool
	Percentage *value.Percentage
	Amount     *value.Money
}

func NewPercentageDiscount(orderID uuid.UUID, percentage value.Percentage) Discount {
	return Discount{
		OrderID:    orderID,
		Percentage: &percentage,
	}
}

func NewAmountDiscount(orderID uuid.UUID, amount value.Money) Discount {
	return Discount{
		OrderID: orderID,
		Amount:  &amount,
	}
}

func (discount Discount) Apply(base value.Money) value.Money {
	if discount.Percentage != nil {
		return base.MultiplyPercentage(*discount.Percentage)
	} else if discount.Amount != nil {
		return *discount.Amount
	}
	return value.Money{}
}
