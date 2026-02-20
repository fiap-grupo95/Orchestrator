package main

import (
	"fmt"
	"github.com/daniloAleite/orchestrator/internal/infrastructure/observability"
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
	newRelicApp, err := observability.NewRelicApp()
	if err != nil {
		logs.Logger().Error().Err(err).Msg("Failed to initialize New Relic; continuing without New Relic integration")
	}

	// Initialize logger with New Relic integration - Global Variable
	logs.Init(newRelicApp)

	hc := httpclient.New()

	osClient := clients.NewOSClient(cfg.OSBaseURL, cfg.OSAuthToken, hc)
	billingClient := clients.NewBillingClient(cfg.BillingBaseURL, hc)
	execClient := clients.NewExecutionClient(cfg.ExecutionBaseURL, hc)
	entityClient := clients.NewEntityAPIClient(cfg.EntityBaseURL, hc)

	uc := usecase.NewOrchestrateServiceOrder(osClient, billingClient, execClient)
	ucCancel := usecase.NewCancelOSUseCase(osClient, entityClient, execClient, billingClient)

	h := handlers.NewOrchestrationHandler(uc, ucCancel)

	mux := http.NewServeMux()
	routes.Register(mux, h)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	logs.Info(fmt.Sprintf("orchestrator running on port %s", cfg.Port))
	_ = srv.ListenAndServe()
}
