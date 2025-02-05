package entities

import "time"

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
