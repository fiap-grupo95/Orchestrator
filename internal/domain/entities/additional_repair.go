package entities

import (
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
	"time"
)

type AdditionalRepair struct {
	ID             string                             `json:"id"`
	Description    string                             `json:"description"`
	ServiceOrderID string                             `json:"service_order_id"`
	ServiceOrder   *ServiceOrder                      `json:"service_order,omitempty"`
	Status         valueobject.AdditionalRepairStatus `json:"status,omitempty"`
	Estimate       *Estimate                          `json:"estimate,omitempty"`
	CreatedAt      time.Time                          `json:"created_at"`
	UpdatedAt      time.Time                          `json:"updated_at"`
	PartsSupplies  []PartsSupply                      `json:"parts_supplies,omitempty"`
	Services       []Service                          `json:"services,omitempty"`
}
