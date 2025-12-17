package logger

import (
	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("failed to create zap logger: " + err.Error())
	}
	return logger
}
