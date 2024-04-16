package models

import (
	"testing"
)

type testpair struct {
	ClientId int
	Amount   float64
}

var tests = []testpair{
	{ClientId: 1, Amount: 0.04},
	{ClientId: 2, Amount: 0},
}

var FakeClientCommissions = []ClientCommission{
	{ClientID: 1, Amount: 0.04},
}

func TestGetCommissionByClientID(t *testing.T) {
	for _, pair := range tests {
		got := GetCommissionByClientID(pair.ClientId, FakeClientCommissions)
		want := pair.Amount

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
