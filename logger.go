package principal

import (
	"github.com/jimenezmaximiliano/kirk"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// fullLogger represents the logger that will be returned by NewLogger.
// We are not using kirk.Logger so a breaking change in kirk, becomes an explicit error here.
type fullLogger interface {
	Error(err error)
	Panic(err error)
	Debug(message string)
	Info(message string)
	Warn(message string)
}

// Check that kirk.ZapLoggerAdapter implements fullLogger.
var _ fullLogger = kirk.ZapLoggerAdapter{}

// NewZapLogger uses kirk to create a new zap sugared logger.
func NewZapLogger() (kirk.ZapLoggerAdapter, error) {
	baseLogger, err := zap.NewProduction()
	if err != nil {
		return kirk.ZapLoggerAdapter{}, errors.Wrap(err, "failed to create a base zap logger")
	}

	sugaredLogger := baseLogger.Sugar()
	if sugaredLogger == nil {
		return kirk.ZapLoggerAdapter{}, errors.Wrap(err, "failed to create a zap sugared logger")
	}

	return kirk.NewLoggerFromSugaredZap(*sugaredLogger), nil
}
