package logger

import (
	"log/slog"
)

// Logger interface
type ILogger interface {
	Info(msg any, keyvals ...any)
	Warn(msg any, keyvals ...any)
	Error(msg any, keyvals ...any)
	Debug(msg any, keyvals ...any)
	Active() bool
	SetActive(active bool)
}

// Logger struct
type Logger struct {
	settings Settings
}

// Settings struct
type Settings struct {
	active bool
}

// Logger global variable
var LOG *Logger

// Initialize logger
func init() {
	LOG = NewLogger()
	LOG.Info("Logger initialized")
}

// Return an instance of Logger struct
func NewLogger() *Logger {
	return &Logger{
		settings: Settings{
			active: true,
		},
	}
}

// Info logs a message at the info level
func (l *Logger) Info(msg string, keyvals ...any) {
	slog.Info(msg, keyvals...)
}

// Warn logs a message at the warn level
func (l *Logger) Warn(msg string, keyvals ...any) {
	slog.Warn(msg, keyvals...)
}

// Error logs a message at the error level
func (l *Logger) Error(msg string, keyvals ...any) {
	slog.Error(msg, keyvals...)
}

// Debug logs a message at the debug level
func (l *Logger) Debug(msg string, keyvals ...any) {
	slog.Debug(msg, keyvals...)
}

// Active returns the logger active status
func (l *Logger) Active() bool {
	return l.settings.active
}

// SetActive sets the logger active
func (l *Logger) SetActive(active bool) {
	l.Warn("Logger active status changed", "active", active)
	l.settings.active = active
}
