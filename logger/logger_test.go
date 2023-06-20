package logger

import "testing"

func TestLogger(t *testing.T) {
	t.Log(Init("debug"))
	Info("hello world", "hello world", true, true, true, 1, 2, 3)
	Debug("hello world", "hello world", true, true, true, 1, 2, 3)
}
