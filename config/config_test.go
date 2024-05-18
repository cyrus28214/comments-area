package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := GetConfig("../config.yaml")
	t.Logf("Config: %+v", cfg)
}
