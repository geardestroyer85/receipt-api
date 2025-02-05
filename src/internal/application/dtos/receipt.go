package dtos

type ItemDto struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessReceiptRequestDto struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []ItemDto `json:"items"`
	Total        string    `json:"total"`
}

type ProcessResponseDto struct {
	ID string `json:"id"`
}

type GetPointsResponseDto struct {
	Points int `json:"points"`
}
