package entities

import (
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
	"time"
)

type Vehicle struct {
	ID         string            `json:"id"`
	Plate      valueobject.Plate `json:"plate"`
	CustomerID string            `json:"customer_id"`
	Customer   *Customer         `json:"customer,omitempty"`
	Model      string            `json:"model"`
	Year       string            `json:"year"`
	Brand      string            `json:"brand"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  *time.Time        `json:"updated_at,"`
	DeletedAt  *time.Time        `json:"deleted_at,omitempty"`
	//ServiceOrders []ServiceOrder    `json:"service_orders,omitempty"`
}
