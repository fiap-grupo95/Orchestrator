package response

import (
	"github.com/daniloAleite/orchestrator/internal/domain/entities"
	"time"
)

type ServiceOrderResponse struct {
	ID                string                                 `json:"id"`
	CustomerID        string                                 `json:"customer_id,omitempty"`
	Customer          *CustomerResponse                      `json:"customer,omitempty"`
	VehicleID         string                                 `json:"vehicle_id,omitempty"`
	Vehicle           *VehicleResponse                       `json:"vehicle,omitempty"`
	Status            string                                 `json:"status"`
	Estimate          *EstimateResponse                      `json:"estimate,omitempty"`
	Execution         *ExecutionResponse                     `json:"execution,omitempty"`
	CreatedAt         *time.Time                             `json:"created_at,omitempty"`
	UpdatedAt         *time.Time                             `json:"updated_at,omitempty"`
	Services          []ServiceOrderServiceResponse          `json:"services,omitempty"`
	PartsSupplies     []ServiceOrderPartsSupplyResponse      `json:"parts_supplies,omitempty"`
	AdditionalRepairs []ServiceOrderAdditionalRepairResponse `json:"additional_repairs,omitempty"`
	PaymentID         *string                                `json:"payment_id,omitempty"`
}

type ServiceOrderServiceResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ServiceOrderPartsSupplyResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ServiceOrderAdditionalRepairResponse struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Estimate    *EstimateResponse `json:"estimate,omitempty"`
}

type ServiceOrderPaymentResponse struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
}

func NewServiceOrderResponse(entity *entities.ServiceOrder) ServiceOrderResponse {
	if entity == nil {
		return ServiceOrderResponse{}
	}

	response := ServiceOrderResponse{
		ID:                entity.ID,
		CustomerID:        entity.CustomerID,
		Customer:          mapCustomerResponse(entity.Customer),
		VehicleID:         entity.VehicleID,
		Vehicle:           mapVehicleResponse(entity.Vehicle),
		Status:            entity.Status.String(),
		CreatedAt:         entity.CreatedAt,
		UpdatedAt:         entity.UpdatedAt,
		Services:          mapServiceResponses(entity.Services),
		PartsSupplies:     mapPartsSupplyResponses(entity.PartsSupplies),
		AdditionalRepairs: mapAdditionalRepairResponses(entity.AdditionalRepairs),
	}

	if entity.Estimate != nil {
		response.Estimate = &EstimateResponse{
			ID:                 entity.Estimate.ID,
			ServiceOrderID:     entity.Estimate.ServiceOrderID,
			Value:              entity.Estimate.Value,
			Status:             string(entity.Estimate.Status),
			AdditionalRepairID: entity.Estimate.AdditionalRepairID,
		}
	}

	if entity.Execution != nil {
		response.Execution = &ExecutionResponse{
			ID:             entity.Execution.ID,
			ServiceOrderID: entity.Execution.ServiceOrderID,
			Status:         string(entity.Execution.Status),
			StartedAt:      entity.Execution.StartedAt,
			FinishedAt:     entity.Execution.FinishedAt,
		}
	}

	return response
}

func NewServiceOrderListResponse(orders []*entities.ServiceOrder) []ServiceOrderResponse {
	if len(orders) == 0 {
		return []ServiceOrderResponse{}
	}

	result := make([]ServiceOrderResponse, 0, len(orders))
	for _, order := range orders {
		result = append(result, NewServiceOrderResponse(order))
	}
	return result
}

func mapServiceResponses(services []entities.Service) []ServiceOrderServiceResponse {
	if len(services) == 0 {
		return nil
	}

	result := make([]ServiceOrderServiceResponse, 0, len(services))
	for _, service := range services {
		result = append(result, ServiceOrderServiceResponse{
			ID:    service.ID,
			Name:  service.Name,
			Price: service.Price,
		})
	}
	return result
}

func mapPartsSupplyResponses(partsSupplies []entities.PartsSupply) []ServiceOrderPartsSupplyResponse {
	if len(partsSupplies) == 0 {
		return nil
	}

	result := make([]ServiceOrderPartsSupplyResponse, 0, len(partsSupplies))
	for _, ps := range partsSupplies {
		result = append(result, ServiceOrderPartsSupplyResponse{
			ID:       ps.ID,
			Price:    ps.Price,
			Quantity: ps.Quantity,
		})
	}
	return result
}

func mapAdditionalRepairResponses(repairs []entities.AdditionalRepair) []ServiceOrderAdditionalRepairResponse {
	if len(repairs) == 0 {
		return nil
	}

	result := make([]ServiceOrderAdditionalRepairResponse, 0, len(repairs))
	for _, repair := range repairs {
		s := ServiceOrderAdditionalRepairResponse{
			ID:          repair.ID,
			Description: repair.Description,
			Status:      repair.Status.String(),
		}

		if repair.Estimate != nil {
			s.Estimate = &EstimateResponse{
				ID:                 repair.Estimate.ID,
				ServiceOrderID:     repair.ServiceOrderID,
				AdditionalRepairID: repair.Estimate.AdditionalRepairID,
				Value:              repair.Estimate.Value,
				Status:             repair.Estimate.Status,
			}
		}
		result = append(result, s)
	}
	return result
}

func mapCustomerResponse(customer *entities.Customer) *CustomerResponse {
	if customer == nil {
		return nil
	}

	return &CustomerResponse{
		ID:          customer.ID,
		UserID:      customer.UserID,
		FullName:    customer.FullName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Document:    customer.CpfCnpj.String(),
	}
}

func mapVehicleResponse(vehicle *entities.Vehicle) *VehicleResponse {
	if vehicle == nil {
		return nil
	}

	return &VehicleResponse{
		ID:         vehicle.ID,
		CustomerID: vehicle.CustomerID,
		Brand:      vehicle.Brand,
		Model:      vehicle.Model,
		Year:       vehicle.Year,
		Plate:      vehicle.Plate.String(),
	}
}
