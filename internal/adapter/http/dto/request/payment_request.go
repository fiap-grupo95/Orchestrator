package request

import (
	"errors"
	"fmt"
	"github.com/fiap-grupo95/os-service-api/internal/domain/entities"
	"time"
)

// PaymentCreateRequest represents the payload to create a payment.
type PaymentCreateRequest struct {
	EstimateID  string  `json:"estimate_id" binding:"required"`
	PaymentDate string  `json:"payment_date" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}

var paymentDateLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
}

var errInvalidPaymentDate = errors.New("invalid payment_date format")

// ToEntity converts the request into a Payment entity parsing accepted date layouts.
func (r PaymentCreateRequest) ToEntity() (entities.Payment, error) {
	paymentDate, err := parsePaymentDate(r.PaymentDate)
	if err != nil {
		return entities.Payment{}, err
	}

	return entities.Payment{
		EstimateID:  r.EstimateID,
		PaymentDate: paymentDate,
		Amount:      r.Amount,
	}, nil
}

func parsePaymentDate(value string) (time.Time, error) {
	for _, layout := range paymentDateLayouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("%w: use one of %v", errInvalidPaymentDate, paymentDateLayouts)
}
