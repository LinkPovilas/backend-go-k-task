package handlers

import "github.com/LinkPovilas/backend-go-k-task/models"

type HighTurnover struct {
	next CommissionHandler
}

func (h *HighTurnover) SetNext(ch CommissionHandler) {
	h.next = ch
}

func (h *HighTurnover) Handle(trx *models.Transaction) error {
	turnover, err := models.GetTotalAmountSinceMonthStart(trx.ClientID, trx.Date)
	if err != nil {
		return err
	}

	if turnover >= 1000 {
		trx.CommissionAmount = 0.03
	}

	if h.next != nil {
		return h.next.Handle(trx)
	}

	return nil
}
