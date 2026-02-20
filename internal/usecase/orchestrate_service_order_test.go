package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
)

type orchestrateOSStub struct {
	createOSFn   func(ctx context.Context, in request.StartInput) (string, error)
	cancelOSFn   func(ctx context.Context, osID string) error
	cancelOSCall int
}

func (s *orchestrateOSStub) GetOS(_ context.Context, _ string) (*response.ServiceOrderResponse, error) {
	return nil, nil
}

func (s *orchestrateOSStub) CreateOS(ctx context.Context, in request.StartInput) (string, error) {
	return s.createOSFn(ctx, in)
}

func (s *orchestrateOSStub) CancelOS(ctx context.Context, osID string) error {
	s.cancelOSCall++
	if s.cancelOSFn != nil {
		return s.cancelOSFn(ctx, osID)
	}
	return nil
}

type orchestrateBillingStub struct {
	createBudgetFn   func(ctx context.Context, osID string) (string, error)
	cancelBudgetFn   func(ctx context.Context, budgetID string) error
	cancelBudgetCall int
	lastCancelBudget string
}

func (s *orchestrateBillingStub) CreateBudget(ctx context.Context, osID string) (string, error) {
	return s.createBudgetFn(ctx, osID)
}

func (s *orchestrateBillingStub) CancelBudget(ctx context.Context, budgetID string) error {
	s.cancelBudgetCall++
	s.lastCancelBudget = budgetID
	if s.cancelBudgetFn != nil {
		return s.cancelBudgetFn(ctx, budgetID)
	}
	return nil
}

type orchestrateExecStub struct {
	startExecutionFn func(ctx context.Context, osID string) error
}

func (s *orchestrateExecStub) StartExecution(ctx context.Context, osID string) error {
	return s.startExecutionFn(ctx, osID)
}

func (s *orchestrateExecStub) CancelExecution(_ context.Context, _ string) error {
	return nil
}

func TestOrchestrateServiceOrder_StartSuccess(t *testing.T) {
	osStub := &orchestrateOSStub{
		createOSFn: func(context.Context, request.StartInput) (string, error) { return "os-1", nil },
	}
	billingStub := &orchestrateBillingStub{
		createBudgetFn: func(context.Context, string) (string, error) { return "b-1", nil },
	}
	execStub := &orchestrateExecStub{
		startExecutionFn: func(context.Context, string) error { return nil },
	}

	uc := NewOrchestrateServiceOrder(osStub, billingStub, execStub)
	out, err := uc.Start(context.Background(), request.StartInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Status != "COMPLETED" || out.OSID != "os-1" || out.BudgetID != "b-1" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestOrchestrateServiceOrder_CreateOSFailure(t *testing.T) {
	osStub := &orchestrateOSStub{
		createOSFn: func(context.Context, request.StartInput) (string, error) { return "", errors.New("os error") },
	}
	billingStub := &orchestrateBillingStub{
		createBudgetFn: func(context.Context, string) (string, error) { return "ignored", nil },
	}
	execStub := &orchestrateExecStub{
		startExecutionFn: func(context.Context, string) error { return nil },
	}

	uc := NewOrchestrateServiceOrder(osStub, billingStub, execStub)
	out, err := uc.Start(context.Background(), request.StartInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if out.Status != "FAILED" {
		t.Fatalf("expected FAILED, got %s", out.Status)
	}
}

func TestOrchestrateServiceOrder_CreateBudgetFailureCompensatesOS(t *testing.T) {
	osStub := &orchestrateOSStub{
		createOSFn: func(context.Context, request.StartInput) (string, error) { return "os-2", nil },
	}
	billingStub := &orchestrateBillingStub{
		createBudgetFn: func(context.Context, string) (string, error) { return "", errors.New("billing error") },
	}
	execStub := &orchestrateExecStub{
		startExecutionFn: func(context.Context, string) error { return nil },
	}

	uc := NewOrchestrateServiceOrder(osStub, billingStub, execStub)
	out, err := uc.Start(context.Background(), request.StartInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if out.Status != "COMPENSATED" || out.OSID != "os-2" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if osStub.cancelOSCall != 1 {
		t.Fatalf("expected CancelOS call once, got %d", osStub.cancelOSCall)
	}
}

func TestOrchestrateServiceOrder_StartExecutionFailureCompensatesAll(t *testing.T) {
	osStub := &orchestrateOSStub{
		createOSFn: func(context.Context, request.StartInput) (string, error) { return "os-3", nil },
	}
	billingStub := &orchestrateBillingStub{
		createBudgetFn: func(context.Context, string) (string, error) { return "b-3", nil },
	}
	execStub := &orchestrateExecStub{
		startExecutionFn: func(context.Context, string) error { return errors.New("exec error") },
	}

	uc := NewOrchestrateServiceOrder(osStub, billingStub, execStub)
	out, err := uc.Start(context.Background(), request.StartInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if out.Status != "COMPENSATED" || out.OSID != "os-3" || out.BudgetID != "b-3" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if osStub.cancelOSCall != 1 {
		t.Fatalf("expected CancelOS call once, got %d", osStub.cancelOSCall)
	}
	if billingStub.cancelBudgetCall != 1 || billingStub.lastCancelBudget != "b-3" {
		t.Fatalf("unexpected billing compensation call count=%d budget=%s", billingStub.cancelBudgetCall, billingStub.lastCancelBudget)
	}
}
