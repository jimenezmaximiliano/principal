package principal

import (
	"github.com/jimenezmaximiliano/kirk"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// FullLogger represents the logger that will be returned by NewLogger.
// We are not using kirk.Logger so a breaking change in kirk, becomes an explicit error here.
type FullLogger interface {
	Error(err error)
	Panic(err error)
	Debug(message string)
	Info(message string)
	Warn(message string)
}

// NewLogger uses kirk to create a new zap sugared logger abstracted into a generic interface.
func NewLogger() FullLogger {
	baseLogger, err := zap.NewProduction()
	if err != nil {
		panic(errors.Wrap(err, "failed to create a base zap logger"))
	}
	sugaredLogger := baseLogger.Sugar()
	if sugaredLogger == nil {
		panic(errors.Wrap(err, "failed to create a zap sugared logger"))
	}

	return kirk.NewLoggerFromSugaredZap(*sugaredLogger)
}
