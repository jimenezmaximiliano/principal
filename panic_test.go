package principal_test

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/jimenezmaximiliano/principal"
	"github.com/jimenezmaximiliano/principal/mocks"
)

func TestItDoesNotPanicWithoutAnError(test *testing.T) {
	test.Parallel()

	logger := mocks.NewLoggerForPanic(test)

	principal.PanicOnError(logger, nil)
}

func TestItPanicsWithAnError(test *testing.T) {
	test.Parallel()

	err := errors.New("oops")

	logger := mocks.NewLoggerForPanic(test)
	logger.On("Panic", err).Return()

	principal.PanicOnError(logger, err)
}
