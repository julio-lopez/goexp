package iohelp_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/julio-lopez/goexp/internal/junk/iohelp"
)

func TestWriteFileNonExisting(t *testing.T) {
	t.Parallel()

	targetFile := filepath.Join(t.TempDir(), "testfile")

	writeFile(t, targetFile, false)
}

func TestWriteFileOverwriteExisting(t *testing.T) {
	t.Parallel()

	targetFile := filepath.Join(t.TempDir(), "testfile")
	err := os.WriteFile(targetFile, []byte("existing junk"), 0o600)

	require.NoError(t, err)
	writeFile(t, targetFile, true)
}

func TestWriteFileNoOverwriteExisting(t *testing.T) {
	t.Parallel()

	targetFile := filepath.Join(t.TempDir(), "testfile")
	err := os.WriteFile(targetFile, []byte("existing junk"), 0o600)

	require.NoError(t, err)

	err = iohelp.WriteFile(targetFile, []byte("some content"), false)
	require.Error(t, err)
}

func writeFile(t *testing.T, targetFile string, overwrite bool) {
	t.Helper()

	want := []byte("some content foo")

	t.Log("destination file:", targetFile)

	err := iohelp.WriteFile(targetFile, want, overwrite)
	require.NoError(t, err)

	got, err := os.ReadFile(targetFile)
	require.NoError(t, err)

	require.Equal(t, want, got)
}
