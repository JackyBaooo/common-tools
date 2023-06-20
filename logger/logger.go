package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Init(level string) error {
	log.SetReportCaller(true)
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	default:
		return fmt.Errorf("unknown log level: %s", level)
	}
	log.Infof("InitLog success, log level: %s", level)
	return nil
}
