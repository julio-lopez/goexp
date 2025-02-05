package scripts

import (
	"bytes"
	"log"
	"strings"

	"al.essio.dev/pkg/shellescape"

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
		r[i] = shellescape.Quote(c)
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
