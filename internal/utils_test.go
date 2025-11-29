package internal_test

import (
	"path/filepath"
	"testing"
)

func NewTempStateFile(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	return filepath.Join(dir, "state.json")
}
