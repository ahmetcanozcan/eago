package loggers

import (
	"github.com/ahmetcanozcan/eago/config"
	"github.com/sirupsen/logrus"
)

var (
	defaultLogger *logrus.Logger
)

func initDefaultLogger() {
	defaultLogger = logrus.New()

	var formatter logrus.Formatter
	if config.EagoJSON.IsProduction() {
		formatter = getProductionLoggerFormatter()
	} else {
		formatter = getDevelopmentLoggerFormatter()
	}

	defaultLogger.SetFormatter(formatter)
}

// Default returns default logger
func Default() Logger {
	return defaultLogger
}

// GetLoggerEntry :
func GetLoggerEntry(fields map[string]interface{}) Logger {
	logrusFields := logrus.Fields(fields)
	return defaultLogger.WithFields(logrusFields)
}
