package models

type ClientCommission struct {
	ClientID int
	Amount   float64
}

type ClientCommissions []ClientCommission

func GetCommissionByClientID(clientID int, cc ClientCommissions) float64 {
	for _, c := range cc {
		if c.ClientID == clientID {
			return c.Amount
		}
	}
	return 0
}
