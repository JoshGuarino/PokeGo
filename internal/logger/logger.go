package logger

import (
	"log/slog"
)

// Logger interface
type ILogger interface {
	Info(msg string, keyvals ...any)
	Warn(msg string, keyvals ...any)
	Error(msg string, keyvals ...any)
	Debug(msg string, keyvals ...any)
	Active() bool
	SetActive(active bool)
	Level() slog.Level
	SetLevel(level slog.Level)
}

// Logger struct
type Logger struct {
	level    slog.Level
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
		level: slog.LevelInfo,
		settings: Settings{
			active: true,
		},
	}
}

// Info logs a message at the info level
func (l *Logger) Info(msg string, keyvals ...any) {
	l.SetLevel(slog.LevelInfo)
	slog.Info(msg, keyvals...)
}

// Warn logs a message at the warn level
func (l *Logger) Warn(msg string, keyvals ...any) {
	l.SetLevel(slog.LevelWarn)
	slog.Warn(msg, keyvals...)
}

// Error logs a message at the error level
func (l *Logger) Error(msg string, keyvals ...any) {
	l.SetLevel(slog.LevelError)
	slog.Error(msg, keyvals...)
}

// Debug logs a message at the debug level
func (l *Logger) Debug(msg string, keyvals ...any) {
	l.SetLevel(slog.LevelDebug)
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

// Level returns the logger level
func (l *Logger) Level() slog.Level {
	return l.level
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level slog.Level) {
	l.level = level
}
