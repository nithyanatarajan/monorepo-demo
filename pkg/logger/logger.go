package logger

import (
	"go.uber.org/zap"
)

// Logger defines the interface for logging operations
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

// DefaultLogger implements the Logger interface
type DefaultLogger struct {
	logger *zap.Logger
}

// New creates a new logger instance
func New() Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err) // This is a demo, in production you'd want to handle this more gracefully
	}
	return &DefaultLogger{
		logger: logger,
	}
}

// Info logs an informational message
func (l *DefaultLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Error logs an error message
func (l *DefaultLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// Debug logs a debug message
func (l *DefaultLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}
