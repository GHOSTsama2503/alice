package logerr_test

import (
	"alice/common/logerr"
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := logerr.Error(errors.New("sample error"), "key", "sample value")
	if err == nil {
		t.Fatal("logerr is not returning an error")
	}
}
