package test

import (
	"context"
	"go-api/pkg/infrastructure"

	"github.com/stretchr/testify/mock"
)

// LoggerMock struct
type LoggerMock struct {
	mock.Mock
}

// Debug func
func (l *LoggerMock) Debug(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// Info func
func (l *LoggerMock) Info(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// Warn func
func (l *LoggerMock) Warn(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// Error func
func (l *LoggerMock) Error(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// Fatal func
func (l *LoggerMock) Fatal(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// Panic func
func (l *LoggerMock) Panic(ctx context.Context, args ...interface{}) {
	l.Called(ctx, args)
}

// WithFields func
func (l *LoggerMock) WithFields(fields infrastructure.Fields) infrastructure.Logger {
	l.Called(fields)
	return &LoggerMock{}
}
