package routes

import (
	"net/http"

	"github.com/daniloAleite/orchestrator/internal/adapter/http/handlers"
)

func Register(mux *http.ServeMux, h *handlers.OrchestrationHandler) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /orchestrations/service-orders", h.StartServiceOrderFlow)
}
