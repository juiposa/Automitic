package log

import (
	"log"
	"log/slog"
)

var Logger = log.Default()

func Info(msg string, v ...interface{}) {
	slog.Info(msg, v...)
}

func Warn(msg string, v ...interface{}) {
	slog.Warn(msg, v...)
}

func Error(msg string, v ...interface{}) {
	slog.Error(msg, v...)
}

func Debug(msg string, v ...interface{}) {
	slog.Debug(msg, v...)
}
