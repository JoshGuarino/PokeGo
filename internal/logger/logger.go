package logger

import (
	"github.com/charmbracelet/log"
)

// Logger interface
type ILogger interface {
	Info(msg string, keyvals ...any)
	Warn(msg string, keyvals ...any)
	Error(msg string, keyvals ...any)
	Debug(msg string, keyvals ...any)
	newMessage(msg string, level log.Level, keyvals ...any)
	Messages() []Message
	Clear()
	Active() bool
	SetActive(active bool)
	Level() log.Level
	SetLevel(level log.Level)
}

// Logger struct
type Logger struct {
	logger   *log.Logger
	messages []Message
	settings Settings
}

type Message struct {
	msg     string
	level   log.Level
	keyvals []any
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
		logger:   log.Default(),
		messages: []Message{},
		settings: Settings{
			active: true,
		},
	}
}

// Info logs a message at the info level
func (l *Logger) Info(msg string, keyvals ...any) {
	l.logger.Info(msg, keyvals...)
	l.newMessage(msg, log.InfoLevel, keyvals...)
}

// Warn logs a message at the warn level
func (l *Logger) Warn(msg string, keyvals ...any) {
	l.logger.Warn(msg, keyvals...)
	l.newMessage(msg, log.WarnLevel, keyvals...)
}

// Error logs a message at the error level
func (l *Logger) Error(msg string, keyvals ...any) {
	l.logger.Error(msg, keyvals...)
	l.newMessage(msg, log.ErrorLevel, keyvals...)
}

// Debug logs a message at the debug level
func (l *Logger) Debug(msg string, keyvals ...any) {
	l.logger.Debug(msg, keyvals...)
	l.newMessage(msg, log.DebugLevel, keyvals...)
}

// Messages returns the logger messages
func (l *Logger) Messages() []Message {
	return l.messages
}

// NewMessage adds a new message to the logger
func (l *Logger) newMessage(msg string, level log.Level, keyvals ...any) {
	l.messages = append(l.messages, Message{
		msg:     msg,
		level:   level,
		keyvals: keyvals,
	})
}

// Clear clears the logger messages
func (l *Logger) Clear() {
	l.messages = []Message{}
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
func (l *Logger) Level() log.Level {
	return l.logger.GetLevel()
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level log.Level) {
	l.logger.SetLevel(level)
}
