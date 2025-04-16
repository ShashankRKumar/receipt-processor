package controllers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]int)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}
	id := uuid.New().String()
	points := utils.CalculatePoints(receipt)
	receipts[id] = points
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}
