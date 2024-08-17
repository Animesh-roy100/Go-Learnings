package valueobjects

import (
	"errors"
	"math"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	INR Currency = "INR"
)

type Money struct {
	amount   int64
	currency Currency
}

func NewMoney(amount float64, currency Currency) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	return Money{amount: int64(math.Round(amount * 100)), currency: currency}, nil
}

func (m Money) Amount() float64 {
	return float64(m.amount) / 100
}

func (m Money) Currency() Currency {
	return m.currency
}

func (m Money) Add(m2 Money) (Money, error) {
	if m.currency != m2.currency {
		return Money{}, errors.New("currencies don't match")
	}
	return Money{amount: m.amount + m2.amount, currency: m.currency}, nil
}

func (m Money) Subtract(m2 Money) (Money, error) {
	if m.currency != m2.currency {
		return Money{}, errors.New("currencies don't match")
	}
	result := m.amount - m2.amount
	if result < 0 {
		return Money{}, errors.New("result cannot be negative")
	}

	return Money{amount: m.amount - m2.amount, currency: m.currency}, nil
}
