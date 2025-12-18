package logger

import (
	"go.uber.org/zap"
)

func New() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true

	logger, err := cfg.Build()
	if err != nil {
		panic("failed to create zap logger: " + err.Error())
	}
	return logger
}
