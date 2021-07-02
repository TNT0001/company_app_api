package infrastructure

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"go-api/pkg/shared/wraperror"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func newLogrusLogger(config *LoggerConfig) (Logger, error) {
	logLevel, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, wraperror.WithTrace(err, nil, nil)
	}
	lLogger := &logrus.Logger{}
	lLogger.SetLevel(logLevel)

	switch config.Output {
	case LoggerOutputFile:
		fileHandler := &lumberjack.Logger{
			Filename: config.Location,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		}
		lLogger.SetOutput(fileHandler)
	default:
		lLogger.SetOutput(os.Stdout)
	}

	switch config.Format {
	case LoggerFormatJSON:
		lLogger.SetFormatter(&logrus.JSONFormatter{})
	default:
		lLogger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
		})
	}

	return &logrusLogger{logger: lLogger}, nil
}

func (l *logrusLogger) Debug(ctx context.Context, args ...interface{}) {
	l.logger.Debug(ctx, args)
}

func (l *logrusLogger) Info(ctx context.Context, args ...interface{}) {
	l.logger.Info(ctx, args)
}

func (l *logrusLogger) Warn(ctx context.Context, args ...interface{}) {
	l.logger.Warn(ctx, args)
}

func (l *logrusLogger) Error(ctx context.Context, args ...interface{}) {
	l.logger.Error(ctx, args)
}

func (l *logrusLogger) Fatal(ctx context.Context, args ...interface{}) {
	l.logger.Fatal(ctx, args)
}

func (l *logrusLogger) Panic(ctx context.Context, args ...interface{}) {
	l.logger.Fatal(ctx, args)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) Debug(ctx context.Context, args ...interface{}) {
	l.entry.Debug(ctx, args)
}

func (l *logrusLogEntry) Info(ctx context.Context, args ...interface{}) {
	l.entry.Info(ctx, args)
}

func (l *logrusLogEntry) Warn(ctx context.Context, args ...interface{}) {
	l.entry.Warn(ctx, args)
}

func (l *logrusLogEntry) Error(ctx context.Context, args ...interface{}) {
	l.entry.Error(ctx, args)
}

func (l *logrusLogEntry) Fatal(ctx context.Context, args ...interface{}) {
	l.entry.Fatal(ctx, args)
}

func (l *logrusLogEntry) Panic(ctx context.Context, args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
