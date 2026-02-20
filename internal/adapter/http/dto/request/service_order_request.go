package request

import (
	"github.com/daniloAleite/orchestrator/internal/domain/entities"
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
)

type ServiceOrderCreateRequest struct {
	CustomerID string `json:"customer_id" binding:"required"`
	VehicleID  string `json:"vehicle_id" binding:"required"`
}

type ServiceOrderDiagnosisUpdateRequest struct {
	Services      []ServiceOrderServiceItem     `json:"services" binding:"omitempty,dive"`
	PartsSupplies []ServiceOrderPartsSupplyItem `json:"parts_supplies" binding:"omitempty,dive"`
}

type ServiceOrderEstimateUpdateRequest struct {
	ServiceOrderStatus string                        `json:"service_order_status" binding:"required"`
	Services           []ServiceOrderServiceItem     `json:"services,omitempty"`
	PartsSupplies      []ServiceOrderPartsSupplyItem `json:"parts_supplies,omitempty"`
}

type ServiceOrderExecutionUpdateRequest struct {
	ServiceOrderStatus string `json:"service_order_status" binding:"required"`
}

type ServiceOrderDeliveryUpdateRequest struct {
	ServiceOrderStatus string `json:"service_order_status" binding:"required"`
}

type ServiceOrderServiceItem struct {
	ID string `json:"id" binding:"required"`
}

type ServiceOrderPartsSupplyItem struct {
	ID       string `json:"id" binding:"required"`
	Quantity int    `json:"quantity,omitempty"`
}

func (r ServiceOrderCreateRequest) ToEntity() *entities.ServiceOrder {
	return &entities.ServiceOrder{
		CustomerID: r.CustomerID,
		VehicleID:  r.VehicleID,
	}
}

func (r ServiceOrderDiagnosisUpdateRequest) ToEntity(id string) *entities.ServiceOrder {
	return &entities.ServiceOrder{
		ID:            id,
		Services:      mapServiceItemsToEntities(r.Services),
		PartsSupplies: mapPartsSupplyItemsToEntities(r.PartsSupplies),
	}
}

func (r ServiceOrderEstimateUpdateRequest) ToEntity(id string) *entities.ServiceOrder {
	return &entities.ServiceOrder{
		ID:            id,
		Status:        valueobject.ParseServiceOrderStatus(r.ServiceOrderStatus),
		Services:      mapServiceItemsToEntities(r.Services),
		PartsSupplies: mapPartsSupplyItemsToEntities(r.PartsSupplies),
	}
}

func (r ServiceOrderExecutionUpdateRequest) ToEntity(id string) *entities.ServiceOrder {
	return &entities.ServiceOrder{
		ID:     id,
		Status: valueobject.ParseServiceOrderStatus(r.ServiceOrderStatus),
	}
}

func (r ServiceOrderDeliveryUpdateRequest) ToEntity(id string) *entities.ServiceOrder {
	return &entities.ServiceOrder{
		ID:     id,
		Status: valueobject.ParseServiceOrderStatus(r.ServiceOrderStatus),
	}
}

func mapServiceItemsToEntities(items []ServiceOrderServiceItem) []entities.Service {
	if len(items) == 0 {
		return nil
	}

	result := make([]entities.Service, 0, len(items))
	for _, item := range items {
		result = append(result, entities.Service{ID: item.ID})
	}
	return result
}

func mapPartsSupplyItemsToEntities(items []ServiceOrderPartsSupplyItem) []entities.PartsSupply {
	if len(items) == 0 {
		return nil
	}

	result := make([]entities.PartsSupply, 0, len(items))
	for _, item := range items {
		result = append(result, entities.PartsSupply{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}
	return result
}
