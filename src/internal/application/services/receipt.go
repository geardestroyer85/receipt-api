package services

import (
	"fmt"
	"math"
	"receipt-api/src/internal/application/dtos"
	"receipt-api/src/internal/domain/entities"
	"receipt-api/src/internal/domain/repositories"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

type ReceiptService struct {
	receiptRepo repositories.ReceiptRepository
}

func NewReceiptService(receiptRepo repositories.ReceiptRepository) *ReceiptService {
	return &ReceiptService{
		receiptRepo: receiptRepo,
	}
}

func (s *ReceiptService) ProcessReceipt(req *dtos.ProcessReceiptRequestDto) (*dtos.ProcessResponseDto, error) {
	receipt, err := convertToEntity(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to entity: %w", err)

	}

	if err := s.receiptRepo.Store(receipt); err != nil {
		return nil, fmt.Errorf("failed to store receipt: %w", err)
	}

	return &dtos.ProcessResponseDto{
		ID: receipt.ID,
	}, nil

}

func (s *ReceiptService) GetPoints(id string) (*dtos.GetPointsResponseDto, error) {
	receipt, err := s.receiptRepo.FindByID(id)
	fmt.Println(receipt, err)
	if err != nil {
		return nil, err
	}

	points := receipt.Points
	if points == -1 {
		points = calculatePoints(receipt)
		receipt.Points = points
		if err := s.receiptRepo.Store(receipt); err != nil {
			return nil, fmt.Errorf("failed to store receipt: %w", err)
		}
	}

	return &dtos.GetPointsResponseDto{
		Points: points,
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

	id := uuid.New().String()
	points := -1

	return &entities.Receipt{
		ID:           id,
		Retailer:     receiptDto.Retailer,
		PurchaseTime: purchaseTime,
		Items:        items,
		Total:        total,
		Points:       points,
	}, nil

}

func calculatePoints(receipt *entities.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += len(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, receipt.Retailer))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	const roundDollarPoints = 50
	if math.Mod(receipt.Total, 1.0) == 0 {
		points += roundDollarPoints
	}
	// Rule 3: 25 points if the total is a multiple of 0.25
	const quarterPoints = 25
	if math.Mod(receipt.Total*4, 1.0) == 0 {
		points += quarterPoints
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions
	// Calculate points for items whose trimmed description length is divisible by 3
	// Points awarded are 20% of the item price (rounded up)
	itemPoints := 0
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			itemPoints += int(math.Ceil(item.Price * 0.2))
		}
	}
	points += itemPoints

	// Rule 6: 6 points if the day in the purchase date is odd
	if receipt.PurchaseTime.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if receipt.PurchaseTime.Hour() >= 14 && receipt.PurchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
