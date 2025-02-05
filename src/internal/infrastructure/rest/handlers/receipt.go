package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{}
}

func (h *ReceiptHandler) ProcessReceipt(c *gin.Context) {
	c.JSON(http.StatusOK, "Receipt processed")
}

func (h *ReceiptHandler) GetPoints(c *gin.Context) {
	c.JSON(http.StatusOK, "Points retrieved")
}
