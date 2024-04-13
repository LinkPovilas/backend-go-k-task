package handlers

import (
	"math"

	"github.com/LinkPovilas/backend-go-k-task/models"
	"github.com/LinkPovilas/backend-go-k-task/utils"
)

type ClientDiscount struct {
	next CommissionHandler
}

func (h *ClientDiscount) SetNext(ch CommissionHandler) {
	h.next = ch
}

func (h *ClientDiscount) Handle(trx *models.Transaction) error {
	clientCommission := utils.GetWhitelistedClientCommission(trx.ClientId)

	if clientCommission != 0 {
		trx.CommissionAmount = math.Min(clientCommission, trx.CommissionAmount)
	}

	if h.next != nil {
		return h.next.Handle(trx)
	}

	return nil
}
