package cmd_test

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/waelmahrous/wormhole/cmd"
)

func TestTransfer(t *testing.T) {
	from := t.TempDir()
	to := t.TempDir()
	tempFile, err := os.CreateTemp(from, "")

	if err != nil {
		t.Fatalf("Could not create temp file")
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		src     []string
		dst     string
		want    []string
		wantErr bool
	}{
		{
			name:    "fail on no files to send",
			src:     []string{},
			dst:     "",
			want:    []string{},
			wantErr: true,
		},
		{
			name:    "send one file",
			src:     []string{tempFile.Name()},
			dst:     to,
			want:    []string{filepath.Join(to, filepath.Base(tempFile.Name()))},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := cmd.Transfer(tt.src, tt.dst)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Transfer() failed: %v", gotErr)
				}
				return
			}

			for _, w := range tt.want {
				if !slices.Contains(got, w) {
					t.Fatalf("expected %q to be in got %v", w, got)
				}

				if _, err := os.Stat(w); err != nil {
					t.Fatalf("Transfer() failed: %v", err)
				}
			}

			if tt.wantErr {
				t.Fatal("Transfer() succeeded unexpectedly")
			}
		})
	}
}
