package observability

import (
	"context"
	"errors"
	logs "github.com/daniloAleite/orchestrator/internal/infrastructure/logger"
	"github.com/daniloAleite/orchestrator/pkg/utils/metrics"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type MetricsCollector interface {
	IncrementCounter(ctx context.Context, name string, labels map[string]string) error
}

type noopMetricsCollector struct{}

func (n *noopMetricsCollector) IncrementCounter(ctx context.Context, name string, labels map[string]string) error {
	return nil
}

type newRelicMetricsCollector struct {
	app *newrelic.Application
}

func NewNewRelicMetricsCollector(app *newrelic.Application) MetricsCollector {
	return &newRelicMetricsCollector{
		app: app,
	}
}

func (m *newRelicMetricsCollector) IncrementCounter(ctx context.Context, name string, labels map[string]string) error {
	if m.app == nil {
		return errors.New("newrelic application is nil")
	}

	metricName := metrics.BuildMetricName(name, labels)

	m.app.RecordCustomMetric(metricName, 1.0)
	return nil
}

var defaultMetricsCollector MetricsCollector = &noopMetricsCollector{}

func SetMetricsCollector(collector MetricsCollector) {
	if collector == nil {
		defaultMetricsCollector = &noopMetricsCollector{}
		return
	}
	defaultMetricsCollector = collector
}

func IncrementCounter(ctx context.Context, name string, labels map[string]string) {
	if defaultMetricsCollector == nil {
		return
	}
	if err := defaultMetricsCollector.IncrementCounter(ctx, name, labels); err != nil {
		logger := logs.LoggerWithContext(ctx)
		logger.Debug().Err(err).Str("metric_name", name).Msg("failed to collect metric")
	}
}
