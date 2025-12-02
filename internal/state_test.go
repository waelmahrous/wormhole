package internal_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/waelmahrous/wormhole/internal"
)

func TestSaveWormholeState(t *testing.T) {
	var (
		state   = internal.WormholeState{}
		tempDir = t.TempDir()
	)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		state   internal.WormholeState
		wantErr bool
	}{
		{"error when path is a file", fmt.Sprintf("%s/.extension", tempDir), state, true},
		{"success with directory path", fmt.Sprintf("%s", tempDir), state, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := internal.SaveWormholeState(tt.path, tt.state)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SaveWormholeState() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SaveWormholeState() succeeded unexpectedly")
			}
		})
	}
}

func TestUpdateDestination(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		dest    string
		wantErr bool
	}{
		{
			name:    "error on invalid directory",
			path:    filepath.Join(t.TempDir(), "invalid_directory"),
			dest:    t.TempDir(),
			wantErr: true,
		},
		{
			name:    "error on invalid target",
			path:    t.TempDir(),
			dest:    filepath.Join(t.TempDir(), "invalid_directory"),
			wantErr: true,
		},
		{
			name:    "update target destination in state file",
			path:    t.TempDir(),
			dest:    t.TempDir(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := internal.UpdateDestination(tt.path, tt.dest)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateDestination() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateDestination() succeeded unexpectedly")
			}

			state, err := internal.LoadWormholeState(tt.path)
			if err != nil {
				t.Fatalf("LoadWormholeState() failed after update: %v", err)
			}

			if state.Destination != tt.dest {
				t.Fatalf("Destination not updated: got %q, want %q", state.Destination, tt.dest)
			}
		})
	}
}

func TestLoadWormholeState(t *testing.T) {
	tempDir := t.TempDir()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    internal.WormholeState
		wantErr bool
	}{
		{
			name:    "load correct state",
			path:    tempDir,
			wantErr: false,
			want: internal.WormholeState{
				Destination: tempDir,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internal.UpdateDestination(tt.path, tt.want.Destination)

			got, gotErr := internal.LoadWormholeState(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadWormholeState() failed: %v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("LoadWormholeState() succeeded unexpectedly")
			}

			if got != tt.want {
				t.Errorf("LoadWormholeState() = %v, want %v", got, tt.want)
			}
		})
	}
}
