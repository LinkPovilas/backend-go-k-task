package models

type commission struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

func NewCommission(amount float64, currency string) commission {
	return commission{amount, currency}
}
