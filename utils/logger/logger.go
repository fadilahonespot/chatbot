package logger

import (
	"context"
	"os"

	"github.com/fadilahonespot/library/logres"
)

var logresLog logres.Logres

func NewLogger() {
	loggerWritter := false
	if os.Getenv("LOGGER_LOGS_WRITE") == "true" {
		loggerWritter = true
	}
	config := logres.LogresConfig{
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,
		LocalTime:  true,
		FolderPath: os.Getenv("LOGGER_FOLDER_PATH"),
		LogsWrite:  loggerWritter,
	}
	
	logresLog = logres.SetLogger(config)
}

func Info(ctx context.Context, title string, message ...interface{}) {
	logresLog.Info(ctx, title, message...)
}

func Error(ctx context.Context, title string, message ...interface{}) {
	logresLog.Error(ctx, title, message)
}

func TDR(ctx context.Context, request []byte, response []byte) {
	logresLog.TDR(ctx, request, response)
}
