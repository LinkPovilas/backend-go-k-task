package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", getHealthCheck)
	r.POST("/commissions", calculateCommission)

	return r
}
