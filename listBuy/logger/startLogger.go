package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func StartLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err == nil {
		Logger.Info("Successful start logger!")
	} else {
		Logger.Fatal("Dont start logger...")
	}
}
