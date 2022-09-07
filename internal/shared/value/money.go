package value

import (
	"ddd/internal/shared/exceptoins"
	"math"
)

type Currency string

var (
	USD Currency = "usd"
)

type Money struct {
	Amount   int64
	Currency Currency
}

func (money Money) Add(other Money) Money {
	return Money{
		Amount:   money.Amount + other.Amount,
		Currency: money.Currency,
	}
}

func (money Money) Subtract(other Money) (Money, error) {
	var err error
	newAmount := money.Amount - other.Amount
	if newAmount < 0 {
		err = exceptoins.NagetiveBalance{}
	}
	return Money{
		Amount:   newAmount,
		Currency: money.Currency,
	}, err
}

func (money Money) Divide(divide int64) []Money {
	return []Money{}
}

func (money Money) Multiply(multiply int64) Money {
	newAmount := money.Amount * multiply
	return Money{
		Amount:   newAmount,
		Currency: money.Currency,
	}
}

func (money Money) MultiplyPercentage(percentage Percentage) Money {
	newAmount := int64(math.Round(float64(money.Amount) * percentage.ToFloat()))
	return Money{
		Amount:   newAmount,
		Currency: money.Currency,
	}
}
