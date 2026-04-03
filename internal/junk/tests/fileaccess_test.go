package junk_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileTypeAccessOnUnreadableDirectory(t *testing.T) {
	d0 := t.TempDir()

	d1 := filepath.Join(d0, "d1")
	err := os.MkdirAll(d1, 0o700)

	require.NoError(t, err)

	d1d2 := filepath.Join(d1, "d2")
	err = os.MkdirAll(d1d2, 0o700)

	require.NoError(t, err)

	d1f2 := filepath.Join(d1, "f2")
	err = os.WriteFile(d1f2, []byte{1, 3, 5}, 0o600)

	require.NoError(t, err)

	err = os.Chmod(d1, 0o400)
	require.NoError(t, err)

	t.Cleanup(func() { os.Chmod(d1, 0o700) }) //nolint:errcheck

	entries, err := os.ReadDir(d1)
	require.NoError(t, err)

	for i, e := range entries {
		info, err := e.Info()
		if err != nil {
			t.Log(e.Name(), " got error:", err)
		}

		// note: depending on the OS and FS type, it may be possible to get whether
		// the entry is a directory, a regular file or something else
		t.Logf("entry %d: isDir:%-5t type:%s : info:%v, %v", i, e.IsDir(), e.Type(), info, e)
	}

	dh, err := os.Open(d1)
	require.NoError(t, err)
	t.Cleanup(func() { dh.Close() }) //nolint:errcheck

	names, err := dh.Readdirnames(0)
	require.NoError(t, err)

	t.Log("names:", names)
}
