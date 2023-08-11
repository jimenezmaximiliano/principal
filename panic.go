package principal

import (
	"context"
)

// LoggerForPanic represents the methods of a logger needed to log a panic.
type LoggerForPanic interface {
	Panic(ctx context.Context, err error)
}

// PanicOnError takes an error and a logger, and panics if the error is not nil.
func PanicOnError(ctx context.Context, logger LoggerForPanic, err error) {
	if err == nil {
		return
	}

	logger.Panic(ctx, err)
}
