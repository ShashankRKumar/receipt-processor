package routes

import (
	"receipt-processor/receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/receipts")
	{
		api.POST("/process", controllers.ProcessReceipt)
		api.GET("/:id/points", controllers.GetPoints)
	}
}
