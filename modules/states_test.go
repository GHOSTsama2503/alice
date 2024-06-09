package modules_test

import (
	"alice/modules"
	"testing"

	"gopkg.in/telebot.v3"
)

func TestInitStates(t *testing.T) {
	modules.InitStates()
}

func TestSetState(t *testing.T) {
	modules.InitStates()
	modules.SetState(1, func(ctx telebot.Context) error { return nil })
}

func TestGetState(t *testing.T) {
	modules.InitStates()

	modules.SetState(1, func(ctx telebot.Context) error { return nil })

	_, err := modules.GetState(1)
	if err != nil {
		t.Fatalf("error getting user state: %v", err)
	}
}

func TestDeleteState(t *testing.T) {
	modules.InitStates()

	modules.SetState(1, func(ctx telebot.Context) error { return nil })

	if err := modules.DelState(1); err != nil {
		t.Fatalf("error deleting user state: %v", err)
	}
}
