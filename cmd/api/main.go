package main

import (
	"net/http"
	"time"

	"github.com/daniloAleite/orchestrator/internal/adapter/clients"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/handlers"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/routes"
	"github.com/daniloAleite/orchestrator/internal/infrastructure/config"
	"github.com/daniloAleite/orchestrator/internal/infrastructure/httpclient"
	"github.com/daniloAleite/orchestrator/internal/infrastructure/logger"
	"github.com/daniloAleite/orchestrator/internal/usecase"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	hc := httpclient.New()

	osClient := clients.NewOSClient(cfg.OSBaseURL, hc)
	billingClient := clients.NewBillingClient(cfg.BillingBaseURL, hc)
	execClient := clients.NewExecutionClient(cfg.ExecutionBaseURL, hc)

	uc := usecase.NewOrchestrateServiceOrder(log, osClient, billingClient, execClient)

	h := handlers.NewOrchestrationHandler(uc)

	mux := http.NewServeMux()
	routes.Register(mux, h)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Info("orchestrator running", "port", cfg.Port)
	_ = srv.ListenAndServe()
}
