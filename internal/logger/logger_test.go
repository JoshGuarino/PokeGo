package logger

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
)

var logger ILogger = newLogger()

func TestNewLogger(t *testing.T) {
	logger.SetActive(true)
	assert.IsType(t, &Logger{}, logger, "Expected Logger instance to be returned")
}

func TestInfo(t *testing.T) {
	logger.Clear()
	logger.Info("test", "key", "value")
	assert.Equal(t, log.InfoLevel, logger.Messages()[0].level, "Expected info level message to be logged")
}

func TestWarn(t *testing.T) {
	logger.Clear()
	logger.Warn("test", "key", "value")
	assert.Equal(t, log.WarnLevel, logger.Messages()[0].level, "Expected warn level message to be logged")
}

func TestError(t *testing.T) {
	logger.Clear()
	logger.Error("test", "key", "value")
	assert.Equal(t, log.ErrorLevel, logger.Messages()[0].level, "Expected error level message to be logged")
}

func TestDebug(t *testing.T) {
	logger.Clear()
	logger.Debug("test", "key", "value")
	assert.Equal(t, log.DebugLevel, logger.Messages()[0].level, "Expected debug level message to be logged")
}

func TestMessages(t *testing.T) {
	logger.Clear()
	logger.newMessage("test", log.InfoLevel, "key", "value")
	assert.Equal(t, 1, len(logger.Messages()), "Expected 1 message to be logged")
	assert.IsType(t, []message{}, logger.Messages(), "Expected message to be of type Message")
}

func TestNewMessage(t *testing.T) {
	logger.Clear()
	logger.newMessage("test", log.InfoLevel, "key", "value")
	assert.Equal(t, "test", logger.Messages()[0].msg, "Expected message to be logged")
	assert.Equal(t, "key", logger.Messages()[0].keyvals[0], "Expected key to be logged")
	assert.Equal(t, "value", logger.Messages()[0].keyvals[1], "Expected value to be logged")
}

func TestClear(t *testing.T) {
	logger.Info("test", "key", "value")
	logger.Clear()
	assert.Equal(t, 0, len(logger.Messages()), "Expected no messages to be logged")
}

func TestActive(t *testing.T) {
	logger.Clear()
	assert.Equal(t, true, logger.Active(), "Expected logger to be active")
	assert.IsType(t, true, logger.Active(), "Expected logger to be active")
}

func TestSetActive(t *testing.T) {
	logger.SetActive(false)
	logger.Clear()
	logger.Info("test", "key", "value")
	assert.Equal(t, false, logger.Active(), "Expected logger to be inactive")
	assert.Equal(t, 0, len(logger.Messages()), "Expected no messages to be logged")
}

func TestLevel(t *testing.T) {
	logger.SetLevel(log.InfoLevel)
	assert.Equal(t, log.InfoLevel, logger.Level(), "Expected logger level to be info")
	assert.IsType(t, log.InfoLevel, logger.Level(), "Expected logger to reurn Level type")
}

func TestSetLevel(t *testing.T) {
	logger.SetLevel(log.WarnLevel)
	assert.Equal(t, log.WarnLevel, logger.Level(), "Expected logger level to be warn")
}
