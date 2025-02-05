package dtos

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

type ProcessResponseDto struct {
	ID string
}

type GetPointsResponseDto struct {
	Points int
}
