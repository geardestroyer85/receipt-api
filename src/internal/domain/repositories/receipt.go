package repositories

import (
	"receipt-api/src/internal/domain/entities"
)

type ReceiptRepository interface {
	Store(receipt *entities.Receipt) error
	FindByID(id string) (*entities.Receipt, error)
}
