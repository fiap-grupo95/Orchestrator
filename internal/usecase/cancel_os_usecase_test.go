package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
)

type cancelOSStub struct {
	getOSFn      func(ctx context.Context, id string) (*response.ServiceOrderResponse, error)
	cancelOSFn   func(ctx context.Context, osID string) error
	lastCancelOS string
}

func (s *cancelOSStub) GetOS(ctx context.Context, id string) (*response.ServiceOrderResponse, error) {
	return s.getOSFn(ctx, id)
}

func (s *cancelOSStub) CreateOS(_ context.Context, _ request.StartInput) (string, error) {
	return "", nil
}

func (s *cancelOSStub) CancelOS(ctx context.Context, osID string) error {
	s.lastCancelOS = osID
	if s.cancelOSFn != nil {
		return s.cancelOSFn(ctx, osID)
	}
	return nil
}

type cancelEntityStub struct {
	releaseFn func(ctx context.Context, parts []response.ServiceOrderPartsSupplyResponse) error
}

func (s *cancelEntityStub) ReleasePartsSupply(ctx context.Context, parts []response.ServiceOrderPartsSupplyResponse) error {
	return s.releaseFn(ctx, parts)
}

type cancelExecStub struct {
	cancelFn      func(ctx context.Context, osID string) error
	lastCancelOS  string
	cancelInvoked int
}

func (s *cancelExecStub) StartExecution(_ context.Context, _ string) error {
	return nil
}

func (s *cancelExecStub) CancelExecution(ctx context.Context, osID string) error {
	s.cancelInvoked++
	s.lastCancelOS = osID
	return s.cancelFn(ctx, osID)
}

type cancelBillingStub struct {
	cancelFn         func(ctx context.Context, budgetID string) error
	lastCancelBudget string
}

func (s *cancelBillingStub) CreateBudget(_ context.Context, _ string) (string, error) {
	return "", nil
}

func (s *cancelBillingStub) CancelBudget(ctx context.Context, budgetID string) error {
	s.lastCancelBudget = budgetID
	return s.cancelFn(ctx, budgetID)
}

func TestCancelOSUseCase_CancelServiceOrderSuccessFallbackBudgetID(t *testing.T) {
	osStub := &cancelOSStub{
		getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) {
			return &response.ServiceOrderResponse{ID: "os-1"}, nil
		},
	}
	entityStub := &cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }}
	execStub := &cancelExecStub{cancelFn: func(context.Context, string) error { return nil }}
	billingStub := &cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}

	uc := NewCancelOSUseCase(osStub, entityStub, execStub, billingStub)
	if err := uc.CancelServiceOrder(context.Background(), "os-1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if billingStub.lastCancelBudget != "os-1" {
		t.Fatalf("expected fallback budget id os-1, got %s", billingStub.lastCancelBudget)
	}
}

func TestCancelOSUseCase_CancelServiceOrderSuccessUsesPaymentID(t *testing.T) {
	paymentID := "budget-9"
	osStub := &cancelOSStub{
		getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) {
			return &response.ServiceOrderResponse{ID: "os-9", PaymentID: &paymentID}, nil
		},
	}
	entityStub := &cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }}
	execStub := &cancelExecStub{cancelFn: func(context.Context, string) error { return nil }}
	billingStub := &cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}

	uc := NewCancelOSUseCase(osStub, entityStub, execStub, billingStub)
	if err := uc.CancelServiceOrder(context.Background(), "os-9"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if billingStub.lastCancelBudget != paymentID {
		t.Fatalf("expected budget id %s, got %s", paymentID, billingStub.lastCancelBudget)
	}
}

func TestCancelOSUseCase_CancelServiceOrderErrorsByStep(t *testing.T) {
	tests := []struct {
		name   string
		setup  func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub)
	}{
		{
			name: "get os fails",
			setup: func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub) {
				return &cancelOSStub{getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) { return nil, errors.New("get") }},
					&cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }},
					&cancelExecStub{cancelFn: func(context.Context, string) error { return nil }},
					&cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}
			},
		},
		{
			name: "release fails",
			setup: func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub) {
				return &cancelOSStub{getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) { return &response.ServiceOrderResponse{ID: "x"}, nil }},
					&cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return errors.New("release") }},
					&cancelExecStub{cancelFn: func(context.Context, string) error { return nil }},
					&cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}
			},
		},
		{
			name: "cancel execution fails",
			setup: func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub) {
				return &cancelOSStub{getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) { return &response.ServiceOrderResponse{ID: "x"}, nil }},
					&cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }},
					&cancelExecStub{cancelFn: func(context.Context, string) error { return errors.New("exec") }},
					&cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}
			},
		},
		{
			name: "cancel budget fails",
			setup: func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub) {
				return &cancelOSStub{getOSFn: func(context.Context, string) (*response.ServiceOrderResponse, error) { return &response.ServiceOrderResponse{ID: "x"}, nil }},
					&cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }},
					&cancelExecStub{cancelFn: func(context.Context, string) error { return nil }},
					&cancelBillingStub{cancelFn: func(context.Context, string) error { return errors.New("budget") }}
			},
		},
		{
			name: "cancel os fails",
			setup: func() (*cancelOSStub, *cancelEntityStub, *cancelExecStub, *cancelBillingStub) {
				return &cancelOSStub{
						getOSFn:    func(context.Context, string) (*response.ServiceOrderResponse, error) { return &response.ServiceOrderResponse{ID: "x"}, nil },
						cancelOSFn: func(context.Context, string) error { return errors.New("cancel os") },
					},
					&cancelEntityStub{releaseFn: func(context.Context, []response.ServiceOrderPartsSupplyResponse) error { return nil }},
					&cancelExecStub{cancelFn: func(context.Context, string) error { return nil }},
					&cancelBillingStub{cancelFn: func(context.Context, string) error { return nil }}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osStub, entityStub, execStub, billingStub := tt.setup()
			uc := NewCancelOSUseCase(osStub, entityStub, execStub, billingStub)
			if err := uc.CancelServiceOrder(context.Background(), "id"); err == nil {
				t.Fatal("expected error")
			}
		})
	}
}
