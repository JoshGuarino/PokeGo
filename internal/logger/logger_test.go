package logger

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

var logger ILogger = NewLogger()

func TestNewLogger(t *testing.T) {
	assert.IsType(t, &Logger{}, logger, "Expected Logger instance to be returned")
}

func TestInfo(t *testing.T) {
	logger.Info("test", "key", "value")
	assert.Equal(t, slog.LevelInfo, logger.Level(), "Expected logger level to be info")
}

func TestWarn(t *testing.T) {
	logger.Warn("test", "key", "value")
	assert.Equal(t, slog.LevelWarn, logger.Level(), "Expected logger level to be warn")
}

func TestError(t *testing.T) {
	logger.Error("test", "key", "value")
	assert.Equal(t, slog.LevelError, logger.Level(), "Expected logger level to be error")
}

func TestDebug(t *testing.T) {
	logger.Debug("test", "key", "value")
	assert.Equal(t, slog.LevelDebug, logger.Level(), "Expected logger level to be debug")
}

func TestActive(t *testing.T) {
	assert.Equal(t, true, logger.Active(), "Expected logger to be active")
	assert.IsType(t, true, logger.Active(), "Expected logger to be active")
}

func TestSetActive(t *testing.T) {
	logger.SetActive(false)
	assert.Equal(t, false, logger.Active(), "Expected logger to be inactive")
	logger.SetActive(true)
	assert.Equal(t, true, logger.Active(), "Expected logger to be active")
}

func TestLevel(t *testing.T) {
	logger.SetLevel(slog.LevelInfo)
	assert.Equal(t, slog.LevelInfo, logger.Level(), "Expected logger level to be info")
	assert.IsType(t, slog.LevelInfo, logger.Level(), "Expected logger to reurn Level type")
}

func TestSetLevel(t *testing.T) {
	logger.SetLevel(slog.LevelWarn)
	assert.Equal(t, slog.LevelWarn, logger.Level(), "Expected logger level to be warn")
}
