package scripts

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteScript(t *testing.T) {
	scriptContent := []string{"echo foo | cat"}
	destFile := filepath.Join(t.TempDir(), "testfile")

	d, err := os.Getwd()
	require.NoError(t, err)

	t.Log("current dir:", d)

	err = WriteScript(destFile, scriptContent, true)
	require.NoError(t, err)

	c, err := os.ReadFile(destFile)
	require.NoError(t, err)

	t.Logf("content: %s", c)
}

func TestQuoteStrings(t *testing.T) {
	cases := []struct {
		expected string
		input    []string
	}{
		{
			expected: `''`,
			input:    []string{""},
		},
		{
			input:    []string{"foo"},
			expected: `foo`,
		},
		{
			input:    []string{"foo", "bar"},
			expected: `foo bar`,
		},
		{
			input:    []string{"foo*"},
			expected: `'foo*'`,
		},
		{
			input:    []string{"foo bar"},
			expected: `'foo bar'`,
		},
		{
			input:    []string{"foo'bar"},
			expected: `'foo'\''bar'`,
		},
		{
			input:    []string{"'foo"},
			expected: `\''foo'`,
		},
		{
			input:    []string{"foo", "bar*"},
			expected: `foo 'bar*'`,
		},
		{
			input:    []string{"foo'foo", "bar", "baz'"},
			expected: `'foo'\''foo' bar 'baz'\'`,
		},
		{
			input:    []string{`\`},
			expected: `'\'`,
		},
		{
			input:    []string{"'"},
			expected: `\'`,
		},
		{
			input:    []string{`\'`},
			expected: `'\'\'`,
		},
		{
			input:    []string{"a''b"},
			expected: `'a'"''"'b'`,
		},
		{
			input:    []string{"azAZ09_!%+,-./:@^"},
			expected: `azAZ09_!%+,-./:@^`,
		},
		{
			input:    []string{"foo=bar", "command"},
			expected: `'foo=bar' command`,
		},
		{
			input:    []string{"foo=bar", "baz=quux", "command"},
			expected: `'foo=bar' 'baz=quux' command`,
		},
	}

	for i, c := range cases {
		c1 := c
		t.Run("case: "+strconv.Itoa(i), func(t *testing.T) {
			got := quoteStrings(c1.input)
			t.Log("got:", got)
			t.Log("exp:", c.expected)
		})
	}
}
