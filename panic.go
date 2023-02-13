package principal

// LoggerForPanic represents the methods of a logger needed to log a panic.
type LoggerForPanic interface {
	Panic(err error)
}

// PanicOnError takes an error and a logger, and panics if the error is not nil.
func PanicOnError(logger LoggerForPanic, err error) {
	if err == nil {
		return
	}

	logger.Panic(err)
}
