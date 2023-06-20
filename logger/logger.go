package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

func Init(level string) error {
	logrus.SetReportCaller(true)
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		return fmt.Errorf("unknown log level: %s", level)
	}
	log.Printf("InitLog success, log level: %s", level)
	return nil
}
