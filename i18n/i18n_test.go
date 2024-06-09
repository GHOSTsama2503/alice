package i18n

import (
	"testing"
)

func TestInit(t *testing.T) {
	if err := Init("../locales"); err != nil {
		t.Fatal(err)
	}
}
