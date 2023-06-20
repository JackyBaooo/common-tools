package config

import "testing"

func TestInitConfig(t *testing.T) {
	t.Log(Init("../config.json"))
}
