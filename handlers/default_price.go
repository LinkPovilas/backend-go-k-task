package handlers

import (
	"math"

	"github.com/LinkPovilas/backend-go-k-task/models"
)

type DefaultPrice struct {
	next CommissionHandler
}

func (h *DefaultPrice) SetNext(ch CommissionHandler) {
	h.next = ch
}

func (h *DefaultPrice) Handle(trx *models.Transaction) error {
	calculatedFee := trx.Amount * 0.5 / 100
	commission := math.Max(calculatedFee, 0.05)

	trx.CommissionAmount = commission

	if h.next != nil {
		return h.next.Handle(trx)
	}

	return nil

}
