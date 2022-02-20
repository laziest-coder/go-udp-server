package factory

import (
	"github.com/Express-24/courier-location-tracker/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	loggerInstance *log.Logger
)

func GetLogger() *log.Logger {
	if loggerInstance == nil {
		Lock.Lock()
		defer Lock.Unlock()

		conf := config.GetConfig()
		logger := log.New()
		logger.SetOutput(os.Stdout)
		logger.SetLevel(getLevel(conf.LogLevel))

		if strings.ToLower(conf.App.Env) == "dev" {
			logger.SetReportCaller(true)
			logger.SetFormatter(&log.TextFormatter{
				DisableColors:   true,
				TimestampFormat: "2006-01-02 15:04:05",
			})
		} else {
			logger.SetFormatter(&log.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			})
		}

		loggerInstance = logger
	}

	return loggerInstance
}

func getLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "error":
		return log.ErrorLevel
	case "warn":
		return log.WarnLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	default:
		return log.InfoLevel
	}
}
