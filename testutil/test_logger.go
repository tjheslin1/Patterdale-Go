package testutil

import (
	"bytes"
	"log"
)

// TestLogger wraps a log.Logger to expose the log output for testing.
type TestLogger struct {
	Logger    *log.Logger
	logBuffer bytes.Buffer
}

// NewTestLogger intialises a new `TestLogger`.
func NewTestLogger() *TestLogger {
	testLogger := TestLogger{}

	testLogger.Logger = log.New(nil, "", 0)
	testLogger.Logger.SetOutput(&testLogger.logBuffer)

	return &testLogger
}

// LogOutput returns the captured logs as a string.
func (logger *TestLogger) LogOutput() string {
	return logger.logBuffer.String()
}
