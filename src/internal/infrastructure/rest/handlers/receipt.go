package handlers

import (
	"net/http"
	"receipt-api/src/internal/application/dtos"
	"receipt-api/src/internal/application/services"
	"receipt-api/src/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	receiptService *services.ReceiptService
}

func NewReceiptHandler(receiptService *services.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{
		receiptService: receiptService,
	}
}

func (h *ReceiptHandler) ProcessReceipt(c *gin.Context) {
	var req dtos.ProcessReceiptRequestDto

	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid.", http.StatusBadRequest)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	res, err := h.receiptService.ProcessReceipt(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	res, err := h.receiptService.GetPoints(id)

	if err != nil {
		appErr := errors.NewAppError(errors.ErrReceiptNotFound, "No Receipt found for that ID.", http.StatusNotFound)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
