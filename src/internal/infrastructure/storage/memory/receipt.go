package memory

import (
	"context"
	"errors"
	"receipt-api/src/internal/domain/entities"
	"receipt-api/src/internal/domain/repositories"
	"sync"
)

type memoryReceiptRepository struct {
	mutex    sync.RWMutex
	receipts map[string]*entities.Receipt
}

func NewMemoryReceiptRepository() repositories.ReceiptRepository {
	return &memoryReceiptRepository{
		receipts: make(map[string]*entities.Receipt),
	}
}

func (r *memoryReceiptRepository) Store(ctx context.Context, receipt *entities.Receipt) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.receipts[receipt.ID] = receipt
	return nil
}

func (r *memoryReceiptRepository) FindByID(ctx context.Context, id string) (*entities.Receipt, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	receipt, exists := r.receipts[id]
	if !exists {
		return nil, errors.New("receipt not found")
	}
	return receipt, nil
}
