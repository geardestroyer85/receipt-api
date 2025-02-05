package validators

import (
	"receipt-api/src/internal/application/dtos"
	"receipt-api/src/internal/domain/errors"
	"strconv"
	"time"
)

type ReceiptValidator struct{}

func NewReceiptValidator() *ReceiptValidator {
	return &ReceiptValidator{}
}

func (v *ReceiptValidator) ValidateProcessReceiptRequest(req *dtos.ProcessReceiptRequestDto) error {
	if err := v.validateRetailer(req.Retailer); err != nil {
		return err
	}

	if err := v.validateDateTime(req.PurchaseDate, req.PurchaseTime); err != nil {
		return err
	}

	if err := v.validateItems(req.Items); err != nil {
		return err
	}

	if err := v.validateTotal(req.Total); err != nil {
		return err
	}

	return nil
}

func (v *ReceiptValidator) validateRetailer(retailer string) error {
	if len(retailer) == 0 {
		return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
	}
	return nil
}

func (v *ReceiptValidator) validateDateTime(dateStr, timeStr string) error {
	_, err := time.Parse("2006-01-02 15:04", dateStr+" "+timeStr)
	if err != nil {
		return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
	}
	return nil
}

func (v *ReceiptValidator) validateItems(items []dtos.ItemDto) error {
	if len(items) == 0 {
		return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
	}

	for _, item := range items {
		if len(item.ShortDescription) == 0 {
			return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
		}

		if _, err := strconv.ParseFloat(item.Price, 64); err != nil {
			return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
		}
	}

	return nil
}

func (v *ReceiptValidator) validateTotal(total string) error {
	if _, err := strconv.ParseFloat(total, 64); err != nil {
		return errors.NewAppError(errors.ErrInvalidReceiptRequest, "The receipt is invalid", 400)
	}
	return nil
}
