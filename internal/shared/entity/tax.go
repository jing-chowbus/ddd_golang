package entity

import "ddd/internal/shared/value"

type Tax struct {
	TaxRate value.Percentage
	Amount  value.Money
}
