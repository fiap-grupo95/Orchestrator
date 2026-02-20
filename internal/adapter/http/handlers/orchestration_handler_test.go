package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
	"github.com/daniloAleite/orchestrator/internal/usecase"
)

type handlerOSStub struct {
	createOSFn func(ctx context.Context, in request.StartInput) (string, error)
	getOSFn    func(ctx context.Context, id string) (*response.ServiceOrderResponse, error)
	cancelOSFn func(ctx context.Context, osID string) error
}

func (s *handlerOSStub) GetOS(ctx context.Context, id string) (*response.ServiceOrderResponse, error) {
	return s.getOSFn(ctx, id)
}
func (s *handlerOSStub) CreateOS(ctx context.Context, in request.StartInput) (string, error) {
	return s.createOSFn(ctx, in)
}
func (s *handlerOSStub) CancelOS(ctx context.Context, osID string) error {
	return s.cancelOSFn(ctx, osID)
}

type handlerBillingStub struct {
	createBudgetFn func(ctx context.Context, osID string) (string, error)
	cancelBudgetFn func(ctx context.Context, budgetID string) error
}

func (s *handlerBillingStub) CreateBudget(ctx context.Context, osID string) (string, error) {
	return s.createBudgetFn(ctx, osID)
}
func (s *handlerBillingStub) CancelBudget(ctx context.Context, budgetID string) error {
	return s.cancelBudgetFn(ctx, budgetID)
}

type handlerExecStub struct {
	startFn  func(ctx context.Context, osID string) error
	cancelFn func(ctx context.Context, osID string) error
}

func (s *handlerExecStub) StartExecution(ctx context.Context, osID string) error {
	return s.startFn(ctx, osID)
}
func (s *handlerExecStub) CancelExecution(ctx context.Context, osID string) error {
	return s.cancelFn(ctx, osID)
}

type handlerEntityStub struct {
	releaseFn func(ctx context.Context, parts []response.ServiceOrderPartsSupplyResponse) error
}

func (s *handlerEntityStub) ReleasePartsSupply(ctx context.Context, parts []response.ServiceOrderPartsSupplyResponse) error {
	return s.releaseFn(ctx, parts)
}

func newTestHandler(startErr error, cancelErr error) *OrchestrationHandler {
	osStub := &handlerOSStub{
		createOSFn: func(context.Context, request.StartInput) (string, error) {
			if startErr != nil {
				return "", startErr
			}
			return "os-1", nil
		},
		getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) {
			if cancelErr != nil {
				return nil, cancelErr
			}
			return &response.ServiceOrderResponse{ID: "os-1"}, nil
		},
		cancelOSFn: func(context.Context, string) error { return nil },
	}
	billingStub := &handlerBillingStub{
		createBudgetFn: func(context.Context, string) (string, error) {
			if startErr != nil {
				return "", startErr
			}
			return "b-1", nil
		},
		cancelBudgetFn: func(context.Context, string) error { return nil },
	}
	execStub := &handlerExecStub{
		startFn: func(context.Context, string) error { return startErr },
		cancelFn: func(context.Context, string) error {
			return cancelErr
		},
	}
	entityStub := &handlerEntityStub{
		releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil },
	}

	uc := usecase.NewOrchestrateServiceOrder(osStub, billingStub, execStub)
	ucCancel := usecase.NewCancelOSUseCase(osStub, entityStub, execStub, billingStub)
	return NewOrchestrationHandler(uc, ucCancel)
}

func TestOrchestrationHandler_StartServiceOrderFlow(t *testing.T) {
	t.Run("invalid json", func(t *testing.T) {
		h := newTestHandler(nil, nil)
		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders", strings.NewReader("{"))
		rec := httptest.NewRecorder()

		h.StartServiceOrderFlow(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rec.Code)
		}
	})

	t.Run("use case error", func(t *testing.T) {
		h := newTestHandler(errors.New("fail"), nil)
		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders", strings.NewReader(`{"customer_id":"1","vehicle_id":"2"}`))
		rec := httptest.NewRecorder()

		h.StartServiceOrderFlow(rec, req)
		if rec.Code != http.StatusBadGateway {
			t.Fatalf("expected 502, got %d", rec.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		h := newTestHandler(nil, nil)
		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders", strings.NewReader(`{"customer_id":"1","vehicle_id":"2"}`))
		rec := httptest.NewRecorder()

		h.StartServiceOrderFlow(rec, req)
		if rec.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", rec.Code)
		}
	})
}

func TestOrchestrationHandler_CancelServiceOrderFlow(t *testing.T) {
	t.Run("missing id", func(t *testing.T) {
		h := newTestHandler(nil, nil)
		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders//cancel", nil)
		rec := httptest.NewRecorder()

		h.CancelServiceOrderFlow(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rec.Code)
		}
	})

	t.Run("success through mux path value", func(t *testing.T) {
		h := newTestHandler(nil, nil)
		mux := http.NewServeMux()
		mux.HandleFunc("POST /orchestrator/v1/service-orders/{id}/cancel", h.CancelServiceOrderFlow)

		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders/abc/cancel", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusAccepted {
			t.Fatalf("expected 202, got %d", rec.Code)
		}
	})

	t.Run("cancel flow error", func(t *testing.T) {
		h := newTestHandler(nil, errors.New("cancel failed"))
		mux := http.NewServeMux()
		mux.HandleFunc("POST /orchestrator/v1/service-orders/{id}/cancel", h.CancelServiceOrderFlow)

		req := httptest.NewRequest(http.MethodPost, "/orchestrator/v1/service-orders/abc/cancel", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadGateway {
			t.Fatalf("expected 502, got %d", rec.Code)
		}
	})
}
