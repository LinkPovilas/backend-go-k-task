package utils

type clientCommissionWhitelist struct {
	ID     int
	Amount float64
}

var clientCommissions = [1]clientCommissionWhitelist{{ID: 42, Amount: 0.05}}

func GetWhitelistedClientCommission(id int) float64 {
	for _, c := range clientCommissions {
		if c.ID == id {
			return c.Amount
		}
	}
	return 0
}
