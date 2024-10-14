package tg_test

import (
	"github.com/ghostsama2503/alice/common/tg"
	"testing"
)

func TestExtractCmdArgs(t *testing.T) {
	emptyArgs := tg.ExtractCmdArgs("/sample")
	if emptyArgs != "" {
		t.Fatalf("error extracting cmd args, result: %s", emptyArgs)
	}

	args := tg.ExtractCmdArgs("/sample foo bar")
	if args != "foo bar" {
		t.Fatalf("error extracting cmd args, result: %s", args)
	}
}
