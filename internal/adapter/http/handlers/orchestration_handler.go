package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/daniloAleite/orchestrator/internal/usecase"
)

type OrchestrationHandler struct {
	uc *usecase.OrchestrateServiceOrder
}

func NewOrchestrationHandler(uc *usecase.OrchestrateServiceOrder) *OrchestrationHandler {
	return &OrchestrationHandler{uc: uc}
}

func (h *OrchestrationHandler) StartServiceOrderFlow(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var in usecase.StartInput
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
