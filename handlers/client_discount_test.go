package handlers

import (
	"testing"

	"github.com/LinkPovilas/backend-go-k-task/models"
)

var fakeClientCommissions = []models.ClientCommission{
	{ClientID: 42, Amount: 0.05},
}

type testpair struct {
	trx                models.Transaction
	expectedCommission float64
}

var tests = [3]testpair{
	{trx: models.Transaction{ClientId: 42, CommissionAmount: 0.25}, expectedCommission: 0.05},
	{trx: models.Transaction{ClientId: 42, CommissionAmount: 0.03}, expectedCommission: 0.03},
	{trx: models.Transaction{ClientId: 1, CommissionAmount: 0.07}, expectedCommission: 0.07},
}

func TestClientDiscountHandle(t *testing.T) {
	for _, pair := range tests {
		clientDiscount := &ClientDiscount{
			ClientCommissions: fakeClientCommissions,
		}
		clientDiscount.Handle(&pair.trx)

		got := pair.trx.CommissionAmount
		want := pair.expectedCommission

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
