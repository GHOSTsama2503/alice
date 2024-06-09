package i18n

import (
	"testing"
)

func TestGetText(t *testing.T) {
	if err := Load("../locales"); err != nil {
		t.Fatal(err)
	}
}
