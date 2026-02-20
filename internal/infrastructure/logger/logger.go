package logs

import (
	"context"
	"io"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/zerologWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

// New relic example
// https://github.com/newrelic/go-agent/blob/v3/integrations/logcontext-v2/zerologWriter/v1.0.5/v3/integrations/logcontext-v2/zerologWriter/example/main.go

var (
	baseWriter io.Writer
	ctxWriter  *zerologWriter.ZerologWriter
	logs       zerolog.Logger
)

func Init(newRelicApp *newrelic.Application) {
	w := zerologWriter.New(os.Stdout, newRelicApp)
	baseWriter = w
	ctxWriter = &w
	logs = zerolog.New(baseWriter)
}

func Logger() *zerolog.Logger {
	return &logs
}

func LoggerWithContext(ctx context.Context) *zerolog.Logger {
	if ctxWriter == nil || ctx == nil {
		return &logs
	}
	txnWriter := ctxWriter.WithContext(ctx)
	logger := logs.Output(txnWriter)
	return &logger
}

func Error(msg string, err error) {
	logs.Error().Err(err).Msg(msg)
}

func Info(msg string) {
	logs.Info().Msg(msg)
}

func Debug(msg string) {
	logs.Debug().Msg(msg)
}

func Warn(msg string) {
	logs.Warn().Msg(msg)
}

func Fatal(msg string, err error) {
	logs.Fatal().Err(err).Msg(msg)
}

func Panic(msg string, err error) {
	logs.Panic().Err(err).Msg(msg)
}
