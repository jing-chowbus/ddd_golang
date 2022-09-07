package entity

import (
	"ddd/internal/shared/value"
	"time"

	"github.com/google/uuid"
)

type LineItem struct {
	ID               uuid.UUID
	Name             string
	Price            value.Money
	Quantity         int64
	SentAt           *time.Time
	Note             string
	Subtatal         value.Money
	PretaxDiscounts  []Discount
	PretaxFees       []Fee
	Tax              Tax
	PosttaxDiscounts []Discount
	PosttaxFees      []Fee
}

func (lineItem LineItem) GetSubtotal() value.Money {
	subtotal := lineItem.Price.Multiply(lineItem.Quantity)
	for _, discount := range lineItem.PretaxDiscounts {
		subtotal, _ = subtotal.Subtract(discount.Apply(subtotal))
	}
	for _, fee := range lineItem.PretaxFees {
		subtotal = subtotal.Add(fee.Apply(subtotal))
	}
	return subtotal
}

func (lineItem LineItem) GetTax() value.Money {
	tax := lineItem.GetSubtotal().MultiplyPercentage(lineItem.Tax.TaxRate)
	return tax
}

func (lineItem LineItem) GetTotal() value.Money {
	subtotal := lineItem.GetSubtotal()
	tax := lineItem.GetTax()
	afterTax := subtotal.Add(tax)
	for _, discount := range lineItem.PosttaxDiscounts {
		afterTax, _ = afterTax.Subtract(discount.Apply(afterTax))
	}
	for _, fee := range lineItem.PosttaxFees {
		afterTax = afterTax.Add(fee.Apply(afterTax))
	}
	return afterTax
}

func (lineItem LineItem) Send(time time.Time) LineItem {
	lineItem.SentAt = &time
	return lineItem
}

func (lineItem LineItem) HasSent() bool {
	return lineItem.SentAt != nil
}
