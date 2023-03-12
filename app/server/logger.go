package server

import (
	"go.uber.org/zap"
)

func NewLogger(debugMode bool) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	if debugMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = logger.Sync()
	}()

	return logger, nil
}
