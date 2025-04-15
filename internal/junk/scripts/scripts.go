package scripts

import (
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/julio-lopez/goexp/internal/junk/iohelp"
)

func WriteScript(filename string, script []string, overwrite bool) error {
	return iohelp.WriteFile(filename, prepareScript(script), overwrite)
}

func quoteStringSlice(s []string) []string {
	if len(s) == 0 {
		return nil
	}

	r := make([]string, len(s))
	for i, c := range s {
		r[i] = shellQuote(c)
	}

	return r
}

func prepareScript(s []string) []byte {
	var b bytes.Buffer

	_, err := b.WriteString("#!/bin/sh\n\n") // header
	assertNoError(err)                       // b.WriteString always returns a nil error

	_, err = b.WriteString(quoteStrings(s))
	assertNoError(err)

	_, err = b.WriteRune('\n') // trailer
	assertNoError(err)

	return b.Bytes()
}

func quoteStrings(s []string) string {
	return strings.Join(quoteStringSlice(s), " ")
}

// fail on unexpected errors
func assertNoError(err error) {
	if err != nil {
		log.Fatalln("unexpected error:", err)
	}
}

var pattern = regexp.MustCompile(`[^\w@%+=:,./-]`)

// Quote returns a shell-escaped version of the string s. The returned value
// is a string that can safely be used as one token in a shell command line.
func shellQuote(s string) string {
	if len(s) == 0 {
		return "''"
	}

	if pattern.MatchString(s) {
		return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
	}

	return s
}
