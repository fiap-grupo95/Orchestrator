package entities

import (
	"time"
)

type Payment struct {
	ID          string    `json:"id"`
	EstimateID  string    `json:"estimate_id"`
	PaymentDate time.Time `json:"payment_date"`
	Amount      float64   `json:"amount"`
}
