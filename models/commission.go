package models

type Commission struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

func NewCommission(amount float64, currency string) Commission {
	return Commission{amount, currency}
}
