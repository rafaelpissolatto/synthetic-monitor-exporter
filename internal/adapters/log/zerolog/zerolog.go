package logadapter

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	Log(level string, msg string)
}

// LogAdapter is the struct that wraps the zerolog.Logger
type LogAdapter struct {
	logger       zerolog.Logger
	DefaultLevel string
}

// NewLogAdapter creates a new LogAdapter
func NewLogAdapter(defaultLogLevel string) *LogAdapter {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &LogAdapter{
		logger:       log.Logger,
		DefaultLevel: defaultLogLevel,
	}
}

// Log logs a message at the specified level or the default level if not specified
func (l *LogAdapter) Log(msg string) {
	level := l.DefaultLevel
	l.logWithLevel(level, msg)
}

// logWithLevel logs a message at the specified level
func (l *LogAdapter) logWithLevel(level string, msg string) {
	switch level {
	case "info":
		l.logger.Info().Msg(msg)
	case "error":
		l.logger.Error().Msg(msg)
	case "fatal":
		l.logger.Fatal().Msg(msg)
	case "debug":
		l.logger.Debug().Msg(msg)
	default:
		l.logger.Info().Msg(msg)
	}
}

// Info logs a message at the info level
func (l *LogAdapter) Info(msg string) {
	l.logWithLevel("info", msg)
}

// Error logs a message at the error level
func (l *LogAdapter) Error(msg string) {
	l.logWithLevel("error", msg)
}

// Fatal logs a message at the fatal level
func (l *LogAdapter) Fatal(msg string) {
	l.logWithLevel("fatal", msg)
}

// Debug logs a message at the debug level
func (l *LogAdapter) Debug(msg string) {
	l.logWithLevel("debug", msg)
}
