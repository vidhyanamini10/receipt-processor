package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor/api/models"
	"receipt-processor/api/services"
	"receipt-processor/api/utils"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	log := utils.GetLogCtx(c)

	if err := c.ShouldBindJSON(&receipt); err != nil {
		log.Errorf("error occured, %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if !services.ValidateReceipt(&receipt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt data"})
		return
	}

	receiptID := services.ProcessReceipt(&receipt)
	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func GetReceiptPoints(c *gin.Context) {
	log := utils.GetLogCtx(c)
	receiptID := c.Param("id")
	points := services.GetReceiptPoints(receiptID)

	if points == -1 {
		log.Errorf("No receipt found for that id")
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
