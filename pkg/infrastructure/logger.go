package infrastructure

import (
	"context"
	"errors"
	"go-api/pkg/shared/wraperror"
)

// Fields type
type Fields map[string]interface{}

const (
	// LogrusLogger instance logger
	LogrusLogger = "Logrus"

	// LoggerOutputStd output type of logger
	LoggerOutputStd = "stdout"
	// LoggerOutputFile output type of logger
	LoggerOutputFile = "file"

	// LoggerFormatText format of logger
	LoggerFormatText = "text"
	// LoggerFormatJSON format of logger
	LoggerFormatJSON = "json"
)

// Logger is our contract for the logger
type Logger interface {
	Debug(ctx context.Context, args ...interface{})

	Info(ctx context.Context, args ...interface{})

	Warn(ctx context.Context, args ...interface{})

	Error(ctx context.Context, args ...interface{})

	Fatal(ctx context.Context, args ...interface{})

	Panic(ctx context.Context, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// LoggerConfig configuration for logger
type LoggerConfig struct {
	// Level: info, debug, error, warning,... . Default: info
	Level string
	// Output: file, stdout. Default: stdout
	Output string
	// Format: json,text,... . Default: text
	Format string
	// Location location of log file
	Location string
}

// NewLogger returns an instance of logger
func NewLogger(config *LoggerConfig, loggerInstance string) (Logger, error) {
	switch loggerInstance {
	case LogrusLogger:
		logger, err := newLogrusLogger(config)
		if err != nil {
			return nil, err
		}
		return logger, nil
	default:
		return nil, wraperror.WithTrace(errors.New("Invalid logger instance"), wraperror.Fields{"loggerInstance": loggerInstance}, nil)
	}
}
