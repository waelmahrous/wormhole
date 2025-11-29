package internal_test

import (
	"testing"

	"github.com/waelmahrous/wormhole/internal"
)

func TestState_SaveAndLoad(t *testing.T) {
	stateFile := NewTempStateFile(t)

	want := internal.WormholeState{Destination: "first"}
	if err := internal.SaveWormholeState(stateFile, want); err != nil {
		t.Fatalf("SaveWormholeState failed: %v", err)
	}

	got, err := internal.LoadWormholeState(stateFile)
	if err != nil {
		t.Fatalf("LoadWormholeState failed: %v", err)
	}

	if got.Destination != want.Destination {
		t.Fatalf("expected destination %q, got %q", want.Destination, got.Destination)
	}
}

