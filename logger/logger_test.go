package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Log(Init("debug"))
	logrus.Info("hello world")
}
