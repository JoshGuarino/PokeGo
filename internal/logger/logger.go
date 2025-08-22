package logger

import (
	"sync"

	"github.com/charmbracelet/log"
)

// Logger interface
type ILogger interface {
	Info(msg string, keyvals ...any)
	Warn(msg string, keyvals ...any)
	Error(msg string, keyvals ...any)
	Debug(msg string, keyvals ...any)
	newMessage(msg string, level log.Level, keyvals ...any)
	Messages() []message
	Clear()
	Active() bool
	SetActive(active bool)
	Level() log.Level
	SetLevel(level log.Level)
}

// Logger struct
type Logger struct {
	logger   *log.Logger
	messages []message
	settings settings
	lock     sync.Mutex
}

// Message struct
type message struct {
	msg     string
	level   log.Level
	keyvals []any
}

// Settings struct
type settings struct {
	active bool
}

// Logger global variable
var LOG *Logger

// Initialize logger
func init() {
	LOG = newLogger()
	LOG.Info("Logger initialized")
}

// Return an instance of Logger struct
func newLogger() *Logger {
	return &Logger{
		logger:   log.Default(),
		messages: []message{},
		settings: settings{
			active: true,
		},
	}
}

func (l *Logger) log(level log.Level, msg string, keyvals ...any) {
	if !l.Active() {
		return
	}
	l.logger.Log(level, msg, keyvals...)
	l.newMessage(msg, level, keyvals...)
}

// Info logs a message at the info level
func (l *Logger) Info(msg string, keyvals ...any) {
	l.log(log.InfoLevel, msg, keyvals...)
}

// Warn logs a message at the warn level
func (l *Logger) Warn(msg string, keyvals ...any) {
	l.log(log.WarnLevel, msg, keyvals...)
}

// Error logs a message at the error level
func (l *Logger) Error(msg string, keyvals ...any) {
	l.log(log.ErrorLevel, msg, keyvals...)
}

// Debug logs a message at the debug level
func (l *Logger) Debug(msg string, keyvals ...any) {
	l.log(log.DebugLevel, msg, keyvals...)
}

// Messages returns the logger messages
func (l *Logger) Messages() []message {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.messages
}

// NewMessage adds a new message to the logger
func (l *Logger) newMessage(msg string, level log.Level, keyvals ...any) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.messages = append(l.messages, message{
		msg:     msg,
		level:   level,
		keyvals: keyvals,
	})
}

// Clear clears the logger messages
func (l *Logger) Clear() {
	l.Warn("Logger messages cleared")
	l.lock.Lock()
	defer l.lock.Unlock()
	l.messages = []message{}
}

// Active returns the logger active status
func (l *Logger) Active() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.settings.active
}

// SetActive sets the logger active
func (l *Logger) SetActive(active bool) {
	l.Warn("Logger active status changed", "active", active)
	l.lock.Lock()
	defer l.lock.Unlock()
	l.settings.active = active
}

// Level returns the logger level
func (l *Logger) Level() log.Level {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.logger.GetLevel()
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level log.Level) {
	l.Warn("Logger level changed", "Level", level)
	l.lock.Lock()
	defer l.lock.Unlock()
	l.logger.SetLevel(level)
}
