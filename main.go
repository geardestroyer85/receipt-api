package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	"unicode"
)

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

func main() {
	fmt.Println(calculatePoints(&testReceipt))
}
