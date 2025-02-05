package services

import (
	"receipt-api/src/internal/application/dtos"
)

type ReceiptService struct {
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{}
}

func (s *ReceiptService) ProcessReceipt(req *dtos.ProcessReceiptRequestDto) (*dtos.ProcessResponseDto, error) {
	return &dtos.ProcessResponseDto{
		ID: "123",
	}, nil
}

func (s *ReceiptService) GetPoints(id string) (*dtos.GetPointsResponseDto, error) {
	return &dtos.GetPointsResponseDto{
		Points: 10,
	}, nil
}
