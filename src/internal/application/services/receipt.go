package services

import (
	"fmt"
	"receipt-api/src/internal/application/dtos"
	"receipt-api/src/internal/domain/entities"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ReceiptService struct {
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{}
}

func (s *ReceiptService) ProcessReceipt(req *dtos.ProcessReceiptRequestDto) (*dtos.ProcessResponseDto, error) {
	receipt, err := convertToEntity(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to entity: %w", err)

	}

	receipt.ID = uuid.New().String()

	return &dtos.ProcessResponseDto{
		ID: receipt.ID,
	}, nil

}

func (s *ReceiptService) GetPoints(id string) (*dtos.GetPointsResponseDto, error) {
	return &dtos.GetPointsResponseDto{
		Points: 10,
	}, nil
}

func convertToEntity(receiptDto *dtos.ProcessReceiptRequestDto) (*entities.Receipt, error) {
	purchaseTime, err := time.Parse("2006-01-02 15:04", receiptDto.PurchaseDate+" "+receiptDto.PurchaseTime)
	if err != nil {
		return nil, fmt.Errorf("invalid purchase date/time format: %w", err)

	}

	total, err := strconv.ParseFloat(receiptDto.Total, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid total amount: %w", err)
	}

	items := make([]entities.Item, len(receiptDto.Items))

	for i, itemDto := range receiptDto.Items {
		price, err := strconv.ParseFloat(itemDto.Price, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price for item %d: %w", i, err)
		}
		items[i] = entities.Item{
			ShortDescription: itemDto.ShortDescription,
			Price:            price,
		}

	}

	return &entities.Receipt{
		Retailer:     receiptDto.Retailer,
		PurchaseTime: purchaseTime,
		Items:        items,
		Total:        total,
	}, nil
}
