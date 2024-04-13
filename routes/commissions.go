package routes

import (
	"log"
	"net/http"

	"github.com/LinkPovilas/backend-go-k-task/handlers"
	"github.com/LinkPovilas/backend-go-k-task/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func calculateCommission(c *gin.Context) {
	var trx models.Transaction
	err := c.ShouldBindJSON(&trx)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body. Check if all fields are provided."})
		return
	}

	clientDiscount := &handlers.ClientDiscount{}

	highTurnover := &handlers.HighTurnover{}
	highTurnover.SetNext(clientDiscount)

	defaultPrice := &handlers.DefaultPrice{}
	defaultPrice.SetNext(highTurnover)

	currency := &handlers.Currency{}
	currency.SetNext(defaultPrice)
	err = currency.Handle(&trx)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not calculate commission."})
		return
	}

	roundedFee, _ := decimal.NewFromFloat(trx.CommissionAmount).Round(2).Float64()
	trx.CommissionAmount = roundedFee
	trx.CommissionCurrency = trx.Currency

	c.JSON(http.StatusOK, models.NewCommission(trx.CommissionAmount, trx.CommissionCurrency))
}
