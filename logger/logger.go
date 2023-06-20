package logger

import (
	"fmt"
	"github.com/JackyBaooo/common-tools/internal"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func Init(level string) error {
	log.SetReportCaller(true)
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	default:
		return fmt.Errorf("unknown log level: %s", level)
	}
	internal.Infof("InitLog success, log level: %s", level)
	return nil
}

func Info(args ...interface{}) {
	n := len(args)
	args = append(args, make([]interface{}, n-1)...)
	for i := n - 1; i > 0; i-- {
		args[i*2] = args[i]
		args[i*2-1] = " "
	}
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Debug(args ...interface{}) {
	n := len(args)
	args = append(args, make([]interface{}, n-1)...)
	for i := n - 1; i > 0; i-- {
		args[i*2] = args[i]
		args[i*2-1] = " "
	}
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
