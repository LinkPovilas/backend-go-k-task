package models

import (
	"testing"
)

type testpair struct {
	ClientID int
	Amount   float64
}

var tests = []testpair{
	{ClientID: 1, Amount: 0.04},
	{ClientID: 2, Amount: 0},
}

var FakeClientCommissions = []ClientCommission{
	{ClientID: 1, Amount: 0.04},
}

func TestGetCommissionByClientID(t *testing.T) {
	for _, pair := range tests {
		got := GetCommissionByClientID(pair.ClientID, FakeClientCommissions)
		want := pair.Amount

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
