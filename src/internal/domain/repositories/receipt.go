package repositories

import (
	"context"
	"receipt-api/src/internal/domain/entities"
)

type ReceiptRepository interface {
	Store(ctx context.Context, receipt *entities.Receipt) error
	FindByID(ctx context.Context, id string) (*entities.Receipt, error)
}
