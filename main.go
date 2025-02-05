package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type ItemDto struct {
	ShortDescription string
	Price            string
}

type ProcessReceiptRequestDto struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []ItemDto
	Total        string
}

var testProcessReceiptRequestDto = ProcessReceiptRequestDto{
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items: []ItemDto{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		}, {
			ShortDescription: "Emils Cheese Pizza",
			Price:            "12.25",
		}, {
			ShortDescription: "Knorr Creamy Chicken",
			Price:            "1.26",
		}, {
			ShortDescription: "Doritos Nacho Cheese",
			Price:            "3.35",
		}, {
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            "12.00",
		},
	},
	Total: "35.35",
}

func convertToEntity(receiptDto *ProcessReceiptRequestDto) (*Receipt, error) {
	purchaseTime, err := time.Parse("2006-01-02 15:04", receiptDto.PurchaseDate+" "+receiptDto.PurchaseTime)
	if err != nil {
		return nil, fmt.Errorf("invalid purchase date/time format: %w", err)
	}

	total, err := strconv.ParseFloat(receiptDto.Total, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid total amount: %w", err)
	}

	items := make([]Item, len(receiptDto.Items))
	for i, itemDto := range receiptDto.Items {
		price, err := strconv.ParseFloat(itemDto.Price, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price for item %d: %w", i, err)
		}
		items[i] = Item{
			ShortDescription: itemDto.ShortDescription,
			Price:            price,
		}
	}

	return &Receipt{
		Retailer:     receiptDto.Retailer,
		PurchaseTime: purchaseTime,
		Items:        items,
		Total:        total,
	}, nil
}

type Item struct {
	ShortDescription string
	Price            float64
}

type Receipt struct {
	ID           string
	Retailer     string
	PurchaseTime time.Time
	Items        []Item
	Total        float64
	Points       int
}

var testReceipt = Receipt{
	Retailer:     "Target",
	PurchaseTime: time.Date(2022, 1, 1, 13, 1, 0, 0, time.UTC),
	Total:        35.35,
	Items: []Item{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            6.49,
		}, {
			ShortDescription: "Emils Cheese Pizza",
			Price:            12.25,
		}, {
			ShortDescription: "Knorr Creamy Chicken",
			Price:            1.26,
		}, {
			ShortDescription: "Doritos Nacho Cheese",
			Price:            3.35,
		}, {
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            12.00,
		},
	},
}

func calculatePoints(receipt *Receipt) int {
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

func (r *Receipt) Equals(other *Receipt) bool {
	if r == nil || other == nil {
		return r == other
	}
	if r.Retailer != other.Retailer || !r.PurchaseTime.Equal(other.PurchaseTime) || r.Total != other.Total || len(r.Items) != len(other.Items) {
		return false
	}
	for i, item := range r.Items {
		if item != other.Items[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(calculatePoints(&testReceipt))
	parsedReceipt, err := convertToEntity(&testProcessReceiptRequestDto)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedReceipt.Equals(&testReceipt))
	fmt.Println(calculatePoints(parsedReceipt))
}
