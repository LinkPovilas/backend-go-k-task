package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/LinkPovilas/backend-go-k-task/models"
)

type Currency struct {
	next CommissionHandler
}

func (h *Currency) SetNext(ch CommissionHandler) {
	h.next = ch
}

func (h *Currency) Handle(trx *models.Transaction) error {
	if trx.Currency != "EUR" {
		exchangeRatesApiUrl := os.Getenv("EXCHANGE_RATE_API_HOSTNAME")
		historicalRatesUrl, err := url.Parse(exchangeRatesApiUrl)
		if err != nil {
			log.Fatal(err)
			return err
		}

		historicalRatesUrl.Path = "/historical"
		q := historicalRatesUrl.Query()
		q.Set("date", "2021-01-01")
		q.Set("access_key", os.Getenv("EXCHANGE_RATE_API_KEY"))
		q.Set("currencies", "EUR")
		q.Set("source", trx.Currency)
		historicalRatesUrl.RawQuery = q.Encode()

		res, err := http.Get(historicalRatesUrl.String())
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}

		var data models.ExchangeRatesData
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
			return err
		}

		quote := data.Quotes[trx.Currency+"EUR"]

		if quote == 0 {
			log.Fatal("Could not get quote")
			return errors.New("Could not get quote")
		}

		trx.Amount = trx.Amount * quote
		trx.Currency = "EUR"
	}

	if h.next != nil {
		return h.next.Handle(trx)
	}

	return nil
}
