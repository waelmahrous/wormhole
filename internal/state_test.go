package internal_test

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/asdine/storm/v3"
	"github.com/waelmahrous/wormhole/internal"
)

var (
	defaultId = "test"
)

func TestInitWormholeStore(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name string
		path string

		wantErr bool
	}{
		{
			name:    "zero: error on empty state directory",
			path:    "",
			wantErr: true,
		},
		{
			name:    "one: success on good directory",
			path:    tempDir,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := internal.Wormhole{
				ID:          defaultId,
				Destination: internal.DefaultDestination,
				StateDir:    tt.path,
			}
			gotErr := w.InitWormholeStore()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("InitWormholeStore() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("InitWormholeStore() succeeded unexpectedly")
			}

			// Expected in good wormholes
			if _, err := os.Stat(filepath.Join(tempDir, internal.StoreName)); err != nil {
				t.Errorf("InitWormholeStore() failed: %v", err)
			} else {
				if db, err := storm.Open(filepath.Join(tt.path, internal.StoreName)); err != nil {
					t.Errorf("InitWormholeStore() failed: %v", err)
				} else {
					var ws internal.Wormhole
					if err := db.One("ID", defaultId, &ws); err != nil {
						t.Fatalf("expected default wormhole, got error: %v", err)
					}

					db.Close()

					// Check that we dont reset
					w.SetDestination(tt.path)
					if err := w.InitWormholeStore(); err != nil {
						t.Errorf("InitWormholeStore() failed: %v", err)
					} else {
						if dest, _ := w.GetDestination(); dest == "" {
							t.Errorf("InitWormholeStore() failed, state got reset: %s", dest)
						}
					}

					db.Close()
				}
			}

		})
	}
}

func TestGetDestination(t *testing.T) {
	tests := []struct {
		name string

		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "zero: error on empty destination",
			path:    t.TempDir(),
			want:    "",
			wantErr: true,
		},
		{
			name:    "one: return correct destination",
			path:    t.TempDir(),
			want:    t.TempDir(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		w := internal.Wormhole{
			ID:          defaultId,
			Destination: internal.DefaultDestination,
			StateDir:    tt.path,
		}
		if err := w.InitWormholeStore(); err != nil {
			t.Fatal(err)
		}

		if _, err := w.SetDestination(tt.want); err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := w.GetDestination()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetDestination() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetDestination() succeeded unexpectedly")
			}

			if tt.want != got {
				t.Errorf("GetDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetDestination(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name string

		path    string
		target  string
		want    internal.Wormhole
		wantErr bool
	}{
		{
			name:    "zero: success set destination",
			path:    t.TempDir(),
			target:  tempDir,
			want:    internal.Wormhole{Destination: tempDir},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		w := internal.Wormhole{
			ID:          defaultId,
			Destination: internal.DefaultDestination,
			StateDir:    tt.path,
		}

		if err := w.InitWormholeStore(); err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := w.SetDestination(tt.target)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SetDestination() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SetDestination() succeeded unexpectedly")
			}

			if tt.want.Destination != got.Destination {
				t.Errorf("SetDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name:    "fail on src is dir",
			src:     []string{from},
			dst:     to,
			want:    []string{},
			wantErr: true,
		},
	}

	w := internal.Wormhole{
		ID:          defaultId,
		Destination: internal.DefaultDestination,
		StateDir:    from,
	}

	if err := w.InitWormholeStore(); err != nil {
		t.Fatal(err)
	}

	if _, err := w.SetDestination(to); err != nil {
		t.Errorf("Transfer() failed: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := internal.TransferRecord{
				Source: tt.src,
				Copy:   false,
			}
			got, gotErr := w.Transfer(record)
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
