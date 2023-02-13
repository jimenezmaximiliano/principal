package principal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// ListenToShutdownSignals returns a context that can be interrupted by syscall.SIGTERM and other signals, and a cancel
// function to cancel the context.
// This allows your app to be shutdown using multiple termination signals.
func ListenToShutdownSignals(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
}
