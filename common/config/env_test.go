package config_test

import (
	"github.com/ghostsama2503/alice/common/config"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	if err := config.LoadEnv(); err != nil {
		t.Fatal("error loading environment")
	}
}
