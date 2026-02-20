package routes

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/handlers"
	"github.com/daniloAleite/orchestrator/internal/usecase"
)

type routeOSStub struct{}

func (s *routeOSStub) GetOS(context.Context, string) (*response.ServiceOrderResponse, error) {
	return &response.ServiceOrderResponse{ID: "x"}, nil
}
func (s *routeOSStub) CreateOS(context.Context, request.StartInput) (string, error) {
	return "x", nil
}
func (s *routeOSStub) CancelOS(context.Context, string) error {
	return nil
}

type routeBillingStub struct{}

func (s *routeBillingStub) CreateBudget(context.Context, string) (string, error) { return "b", nil }
func (s *routeBillingStub) CancelBudget(context.Context, string) error            { return nil }

type routeExecStub struct{}

func (s *routeExecStub) StartExecution(context.Context, string) error  { return nil }
func (s *routeExecStub) CancelExecution(context.Context, string) error { return nil }

type routeEntityStub struct{}

func (s *routeEntityStub) ReleasePartsSupply(context.Context, []response.ServiceOrderPartsSupplyResponse) error {
	return nil
}

func TestRegister(t *testing.T) {
	osStub := &routeOSStub{}
	billingStub := &routeBillingStub{}
	execStub := &routeExecStub{}
	entityStub := &routeEntityStub{}

	h := handlers.NewOrchestrationHandler(
		usecase.NewOrchestrateServiceOrder(osStub, billingStub, execStub),
		usecase.NewCancelOSUseCase(osStub, entityStub, execStub, billingStub),
	)

	mux := http.NewServeMux()
	Register(mux, h)

	healthReq := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthRec := httptest.NewRecorder()
	mux.ServeHTTP(healthRec, healthReq)
	if healthRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", healthRec.Code)
	}

	startReq := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders", strings.NewReader(`{"customer_id":"1","vehicle_id":"2"}`))
	startRec := httptest.NewRecorder()
	mux.ServeHTTP(startRec, startReq)
	if startRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", startRec.Code)
	}

	cancelReq := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders/1/cancel", nil)
	cancelRec := httptest.NewRecorder()
	mux.ServeHTTP(cancelRec, cancelReq)
	if cancelRec.Code != http.StatusAccepted {
		t.Fatalf("expected 202, got %d", cancelRec.Code)
	}
}
