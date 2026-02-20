package observability

import (
	"context"
	logs "github.com/daniloAleite/orchestrator/internal/infrastructure/logger"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApp() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("os-service-api"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		logs.Logger().Error().Err(err).Msg("failed to initialize New Relic")
		return nil, err
	}

	if err := app.WaitForConnection(10 * time.Second); err != nil {
		logs.Logger().Error().Err(err).Msg("failed to wait for new relic connection, the application will continue without Observability")
		return nil, err
	}

	return app, nil
}

func StartSegment(ctx context.Context, name string) func() {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return func() {
			// no-op
		}
	}
	s := txn.StartSegment(name)
	return func() {
		s.End()
	}
}
