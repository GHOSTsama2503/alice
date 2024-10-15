package i18n_test

import (
	"strings"
	"testing"

	"github.com/ghostsama2503/alice/i18n"
)

func TestInit(t *testing.T) {
	if err := i18n.Init("en"); err != nil {
		t.Fatal(err)
	}
}

func TestGetLocale(t *testing.T) {
	locale := i18n.GetLocale("en")

	if locale.LocaleName == "" {
		t.Fatal("empty locale name")
	}
}

func TestGetLocaleDefault(t *testing.T) {
	locale := i18n.GetLocale("")

	if locale.LocaleName == "" {
		t.Fatal("empty locale name")
	}
}

func TestWithOptions(t *testing.T) {
	locale := i18n.GetLocale("")

	opts := i18n.Options{
		"me":   "bot",
		"user": "user",
	}

	text := i18n.WithOptions(locale.StartMessage, opts)

	if strings.ContainsAny(text, "}{") {
		t.Fatal("unexpected result", text)
	}
}
