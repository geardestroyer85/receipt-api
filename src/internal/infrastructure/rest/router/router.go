package router

import (
	"receipt-api/src/internal/infrastructure/rest/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	receiptHandler *handlers.ReceiptHandler
}

func NewRouter(receiptHandler *handlers.ReceiptHandler) *Router {
	return &Router{
		receiptHandler: receiptHandler,
	}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {

	COMMON_API_PATH := ""
	api := engine.Group(COMMON_API_PATH)
	{
		receipt := api.Group("/receipts")
		{
			receipt.POST("/process", r.receiptHandler.ProcessReceipt)
			receipt.GET("/:id/points", r.receiptHandler.GetPoints)
		}

	}
}
