package loggers

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Logger representation
type Logger interface {
	Warn(...interface{})
	Error(...interface{})
	Info(...interface{})
	Fatal(...interface{})
}

func getProductionLoggerFormatter() logrus.Formatter {
	formatter := logrus.JSONFormatter{}
	return &formatter
}

func getDevelopmentLoggerFormatter() logrus.Formatter {
	formatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
	}

	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "red",
		PanicLevelStyle: "red",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "black+h",
	})
	return formatter
}

// InitializeLoggers : Initialize all loggers
func InitializeLoggers() {
	initDebugLogger()
	initDefaultLogger()
}
