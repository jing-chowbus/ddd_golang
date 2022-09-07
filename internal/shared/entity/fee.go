package entity

import (
	"ddd/internal/shared/value"

	"github.com/google/uuid"
)

type Fee struct {
	ID         uuid.UUID
	OrderID    uuid.UUID
	LineItemID uuid.UUID
	Percentage *value.Percentage
	Amount     *value.Money
}

func (fee Fee) Apply(base value.Money) value.Money {
	if fee.Percentage != nil {
		return base.MultiplyPercentage(*fee.Percentage)
	} else if fee.Amount != nil {
		return *fee.Amount
	}
	return value.Money{}
}
