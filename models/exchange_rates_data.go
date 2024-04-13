package models

type ExchangeRatesData struct {
	Success    bool               `json:"success"`
	Terms      string             `json:"terms"`
	Privacy    string             `json:"privacy"`
	Historical bool               `json:"historical"`
	Date       string             `json:"date"`
	Timestamp  int64              `json:"timestamp"`
	Source     string             `json:"source"`
	Quotes     map[string]float64 `json:"quotes"` // key: "source|target" currencies, example: "USDEUR"
}
