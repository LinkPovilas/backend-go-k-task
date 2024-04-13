package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		log.Fatal(err)
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTransactionsTable()
	seedTransactionTable()
}

func createTransactionsTable() {
	createTransactionsTable := `
	CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    date TEXT NOT NULL,
    amount REAL NOT NULL,
    currency TEXT NOT NULL,
    commission_amount REAL NOT NULL,
    commission_currency TEXT NOT NULL);`

	_, err := DB.Exec(createTransactionsTable)

	if err != nil {
		log.Fatal(err)
		panic("could not create events table")
	}
}

func seedTransactionTable() {
	_, err := DB.Exec("DELETE FROM transactions;")
	if err != nil {
		log.Fatal(err)
		panic("could not clear transactions table")
	}

	insertTestData := `
	INSERT INTO transactions (client_id, date, amount, currency, commission_amount, commission_currency) VALUES
	(42, "2021-01-01", 2000.0, "EUR", 0.05, "EUR"),
	(1, "2021-01-03", 500.0, "EUR", 2.5, "EUR"),
	(1, "2021-01-04", 499.0, "EUR", 2.5, "EUR"),
	(1, "2021-01-05", 100.0, "EUR", 0.5, "EUR"),
	(1, "2021-01-06", 1.0, "EUR", 0.03, "EUR"),
	(1, "2021-01-01", 500.0, "EUR", 2.5, "EUR"),
	(7, "2021-01-01", 900.0, "EUR", 4.5, "EUR"),
	(8, "2021-01-01", 3.5, "EUR", 0.05, "EUR"),
	(7, "2021-01-04", 43.33, "EUR", 0.22, "EUR"),
	(10, "2021-01-01", 0.01, "EUR", 0.05, "EUR");`

	_, err = DB.Exec(insertTestData)
	if err != nil {
		log.Fatal(err)
		panic("could not seed transactions table")
	}
}
