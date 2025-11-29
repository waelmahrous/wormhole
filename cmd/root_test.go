package cmd_test

import (
	"strings"
	"testing"

	"github.com/waelmahrous/wormhole/cmd"
)

func TestRootCommand_Help(t *testing.T) {
	out, err := cmd.ExecuteCommand(t, cmd.RootCmd, "--help")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !strings.Contains(out, "Usage:") {
		t.Errorf("expected help output to contain 'Usage:', got:\n%s", out)
	}
}
