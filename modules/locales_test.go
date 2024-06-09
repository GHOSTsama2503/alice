package modules_test

import (
	"alice/common/config"
	"alice/modules"
	"testing"
)

func TestGetUserLocale(t *testing.T) {
	locale := modules.GetUserLocale(0)
	if locale != config.Env.DefaultLocale {
		t.Fatal("unexpected behavior getting user locale")
	}
}

func TestSetUserLocale(t *testing.T) {
	err := modules.SetUserLocale(0, "es")
	if err != nil {
		t.Fatal("unexpected error setting user locale")
	}
}
