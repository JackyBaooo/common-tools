package internal

import "github.com/sirupsen/logrus"

var log = logrus.New()

func Info(args ...interface{}) {
	n := len(args)
	args = append(args, make([]interface{}, n-1)...)
	for i := n - 1; i > 0; i-- {
		args[i*2] = args[i-1]
		args[i*2-1] = " "
	}
	log.Info(args)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
