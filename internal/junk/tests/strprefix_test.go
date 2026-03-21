package junk_test

import (
	"bytes"
	"testing"
)

func BenchmarkStringByteSlicePrefix(b *testing.B) {
	tc := struct {
		in     []byte
		prefix string
		expect bool
	}{
		in:     []byte("this is some text"),
		prefix: "this",
	}
	for b.Loop() {
		_ = hasPrefix(tc.in, tc.prefix)
	}
}

// context: blob prefix comparison in kopia
func hasPrefix(b []byte, prefix string) bool {
	// If the prefix is longer than b, it cannot be a prefix.
	if len(prefix) > len(b) {
		return false
	}

	// Slice the string to the length of the prefix and convert to a byte slice.
	// This conversion does not involve copying the underlying data,
	// but rather creates a new slice header pointing to the same backing array.
	return prefix == string(b[:len(prefix)])
}

func BenchmarkStringByteSlicePrefix2(b *testing.B) {
	tc := struct {
		in     []byte
		prefix string
		expect bool
	}{
		in:     []byte("this is something else"),
		prefix: "this",
	}
	for b.Loop() {
		_ = hasPrefix2(string(tc.in), tc.prefix)
	}
}

func hasPrefix2(b string, prefix string) bool {
	// If the prefix is longer than b, it cannot be a prefix.
	if len(prefix) > len(b) {
		return false
	}

	// Slice the string to the length of the prefix and convert to a byte slice.
	// This conversion does not involve copying the underlying data,
	// but rather creates a new slice header pointing to the same backing array.
	// return prefix == string(b[:len(prefix)])
	// return strings.HasPrefix(string(b), prefix)
	// return strings.HasPrefix(b, prefix)
	return bytes.HasPrefix([]byte(b), []byte(prefix))
}
