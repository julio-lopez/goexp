package mmap_test

import (
	"errors"
	"os"
	"syscall"
	"testing"

	"github.com/edsrzf/mmap-go"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestAppendMmappedSlice(t *testing.T) {
	const (
		mmapFlags = 0
		// segmentSize = 8192
		segmentSize = 1<<20 + 8192 // 2 MB
		// segmentSize = 10
	)

	f := createAutoDeleteTempFile(t)
	require.NotNil(t, f)

	// t.Cleanup(func() {
	// 	require.NoError(t, f.Close())
	// })

	err := f.Truncate(segmentSize)
	require.NoError(t, err)

	s1, err := mmap.MapRegion(f, segmentSize, mmap.RDWR, mmapFlags, 0)
	require.NoError(t, err)
	require.NotNil(t, s1)

	require.NoError(t, f.Close())

	t.Logf("len(s1):%v, cap(s1): %v, &s1:%p, &s1[0]:%p", len(s1), cap(s1), &s1, &s1[0])

	s2 := s1[0 : len(s1)-1]
	s2 = append(s2, 'x')

	t.Logf("len(s2):%v, cap(s2): %v, &s2:%p, &s2[0]:%p", len(s2), cap(s2), &s2, &s2[0])

	// t.Logf("s1: %q", s1)

	// s1 = s1[:0]
	// for i := range cap(s1) {
	// 	s1 = append(s1, byte(i))
	// }

	for i := range s1 {
		s1[i] = byte(i)
	}

	// t.Logf("s1: %q", s1)

	// s2 = append(s2, 0xB)

	t.Logf("len(s2):%v, cap(s2): %v, &s2:%p, &s2[0]:%p", len(s2), cap(s2), &s2, &s2[0])

	// s2[cap(s2)] = 0

	require.Panics(t, func() {
		s2[cap(s2)] = 0
	})
}

func createAutoDeleteTempFile(t *testing.T) *os.File {
	return createAutoDeleteTempFileFallback(t, "/mnt/t")
}

func createAutoDeleteTempFileFallback(t *testing.T, tempDir string) *os.File {
	f, err := os.CreateTemp(tempDir, "kt-")
	require.NoError(t, err)

	t.Log("file name:", f.Name())

	// immediately remove/unlink the file while we keep the handle open.
	err = os.Remove(f.Name())
	require.NoError(t, err)

	return f
}

func TestCreateAutoDeleteTempFile2(t *testing.T) {
	f := createAutoDeleteTempFile2(t)

	require.NoError(t, f.Close())
}

func createAutoDeleteTempFile2(t *testing.T) *os.File {
	const permissions = 0o600

	// tempDir := t.TempDir()
	tempDir := "/mnt/t"

	t.Log("tempdir:", tempDir)

	// const tempFileFlag = unix.O_TMPFILE
	const tempFileFlag = 0
	// on reasonably modern Linux (3.11 and above) O_TMPFILE is supported,
	// which creates invisible, unlinked file in a given directory.
	fd, err := unix.Open(tempDir, unix.O_RDWR|tempFileFlag|unix.O_CLOEXEC, permissions)
	// var err error

	if err == nil {
		return os.NewFile(uintptr(fd), "")
	}

	if errors.Is(err, syscall.EISDIR) || errors.Is(err, syscall.EOPNOTSUPP) {
		return createAutoDeleteTempFileFallback(t, tempDir)
	}

	require.NoError(t, err)

	return nil
}
