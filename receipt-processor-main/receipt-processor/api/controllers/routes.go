package controllers

import "github.com/gin-gonic/gin"

func Initialize(router *gin.Engine) {
	api := router.Group("/receipts")
	{
		api.POST("/process", ProcessReceipt)
		api.GET("/:id/points", GetReceiptPoints)
	}
}
