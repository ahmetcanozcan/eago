package loggers

import (
	"io/ioutil"

	"github.com/ahmetcanozcan/eago/config"
	"github.com/sirupsen/logrus"
)

var (
	// debugLogger is used for debugging process.
	// it's only available if  eagoEnv value is "development" in eago.json
	// otherwise, debugLogger doesn't log anything
	debugLogger *logrus.Logger
)

func initDebugLogger() {
	debugLogger := logrus.New()
	if !config.EagoJSON.IsDevelopment() {
		// Make debug logger silence for non development environment
		debugLogger.SetOutput(ioutil.Discard)
		return
	}
	debugLogger.SetFormatter(getDevelopmentLoggerFormatter())
}

// Debug returns debug logger
func Debug() Logger {
	return debugLogger
}
