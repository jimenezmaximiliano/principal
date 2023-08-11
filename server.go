package principal

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/pseidemann/finish"
)

// Server includes the methods needed from an HTTP server to be used by StartServer.
type Server interface {
	Shutdown(ctx context.Context) error
	ListenAndServe() error
}

// LoggerForStartServer includes the methods needed for StartServer.
type LoggerForStartServer interface {
	Panic(ctx context.Context, err error)
	Error(ctx context.Context, err error)
	Info(ctx context.Context, message string)
}

// StartServer starts an HTTP server that can be shut down gracefully.
func StartServer(ctx context.Context, logger LoggerForStartServer, server Server, shutdownTimeout time.Duration) {
	finisher := &finish.Finisher{
		Timeout: shutdownTimeout,
		Log: finisherLoggerAdapter{
			logger: logger,
		},
		Signals: append(finish.DefaultSignals, syscall.SIGHUP),
	}

	finisher.Add(server)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Panic(ctx, err)
		}
	}()

	finisher.Wait()
}

type finisherLoggerAdapter struct {
	logger LoggerForStartServer
}

// Infof is required for finish.Finisher to log informational messages.
func (logger finisherLoggerAdapter) Infof(format string, v ...interface{}) {
	logger.logger.Info(context.Background(), fmt.Sprintf(format, v...))
}

// Errorf is required for finish.Finisher to log errors.
func (logger finisherLoggerAdapter) Errorf(format string, v ...interface{}) {
	logger.logger.Error(context.Background(), errors.Errorf(format, v...))
}
