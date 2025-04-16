package routes

import (
	"receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/receipts/process", controllers.ProcessReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)
}
