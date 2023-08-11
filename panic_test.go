package principal_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/jimenezmaximiliano/principal"
	"github.com/jimenezmaximiliano/principal/mocks"
)

func TestItDoesNotPanicWithoutAnError(test *testing.T) {
	test.Parallel()

	logger := mocks.NewLoggerForPanic(test)

	principal.PanicOnError(context.Background(), logger, nil)
}

func TestItPanicsWithAnError(test *testing.T) {
	test.Parallel()

	err := errors.New("oops")

	logger := mocks.NewLoggerForPanic(test)
	logger.On("Panic", mock.Anything, err).Return()

	principal.PanicOnError(context.Background(), logger, err)
}
