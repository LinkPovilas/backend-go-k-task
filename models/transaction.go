package models

import (
	"database/sql"
	"fmt"
	"time"

	db "github.com/LinkPovilas/backend-go-k-task/database"
)

type Transaction struct {
	ID                 int64
	ClientId           int     `json:"client_id" binding:"required"`
	Date               string  `json:"date" binding:"required"`
	Amount             float64 `json:"amount" binding:"required"`
	Currency           string  `json:"currency" binding:"required"`
	CommissionAmount   float64
	CommissionCurrency string
}

func (t *Transaction) Save() error {
	query := `INSERT INTO transactions (client_id, date, amount, currency, commission_amount, commission_currency) 
	VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.ClientId, t.Date, t.Amount, t.Currency, t.CommissionAmount, t.CommissionCurrency)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	t.ID = id
	return err
}

func GetTotalAmountSinceMonthStart(clientId int, dateUntil string) (float64, error) {
	query := `
	SELECT SUM(Amount) AS TotalAmount
	FROM transactions
	WHERE client_id = ? AND date BETWEEN ? AND ?;`

	date, err := time.Parse("2006-01-02", dateUntil)
	if err != nil {
		return 0, err
	}

	firstDayOfMonth := fmt.Sprintf("%d-%02d-01", date.Year(), date.Month())
	row := db.DB.QueryRow(query, clientId, firstDayOfMonth, dateUntil)

	var totalAmount sql.NullFloat64
	err = row.Scan(&totalAmount)
	if err != nil {
		return 0, err
	}

	if totalAmount.Valid {
		return totalAmount.Float64, nil
	}

	return 0, nil
}
