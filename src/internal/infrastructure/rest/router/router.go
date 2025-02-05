package router

import (
	"receipt-api/src/internal/infrastructure/rest/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {

	receiptHandler := handlers.NewReceiptHandler()

	COMMON_API_PATH := ""
	api := engine.Group(COMMON_API_PATH)
	{
		receipt := api.Group("/receipts")
		{
			receipt.POST("/process", receiptHandler.ProcessReceipt)
			receipt.GET("/:id/points", receiptHandler.GetPoints)
		}
	}
}
