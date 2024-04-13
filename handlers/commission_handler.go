package handlers

import (
	"github.com/LinkPovilas/backend-go-k-task/models"
)

type CommissionHandler interface {
	SetNext(CommissionHandler)
	Handle(trx *models.Transaction) error
}
