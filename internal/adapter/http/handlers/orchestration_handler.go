package handlers

import (
	"context"
	"encoding/json"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"net/http"
	"time"

	"github.com/daniloAleite/orchestrator/internal/usecase"
)

type OrchestrationHandler struct {
	uc       *usecase.OrchestrateServiceOrder
	ucCancel *usecase.CancelOSUseCase
}

func NewOrchestrationHandler(uc *usecase.OrchestrateServiceOrder, ucCancel *usecase.CancelOSUseCase) *OrchestrationHandler {
	return &OrchestrationHandler{uc: uc, ucCancel: ucCancel}
}

func (h *OrchestrationHandler) StartServiceOrderFlow(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var in request.StartInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	out, err := h.uc.Start(ctx, in)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error":  err.Error(),
			"result": out,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(out)
}

func (h *OrchestrationHandler) CancelServiceOrderFlow(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var in request.StartInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	osID := r.URL.Query().Get("id")
	if osID == "" {
		http.Error(w, "Failed to parse service order ID", http.StatusBadRequest)
		return
	}

	err := h.ucCancel.CancelServiceOrder(ctx, osID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w)
}
