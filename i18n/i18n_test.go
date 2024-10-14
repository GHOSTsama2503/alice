package i18n_test

import (
	"github.com/ghostsama2503/alice/i18n"
	"testing"
)

func TestInit(t *testing.T) {
	if err := i18n.Init(); err != nil {
		t.Fatal(err)
	}
}

func TestGetLocale(t *testing.T) {
	locale, err := i18n.GetLocale("")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(locale.OK)
}
