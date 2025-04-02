package iohelp

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"strings"
)

// WriteFile writes content to a destination file, potentially overwriting the
// destination file.
//
// It returns an error when
// - the file cannot be written;
// - the destination file is an existing directory;
// - the destination file already exists and overwrite is false
func WriteFile(filename string, content []byte, overwrite bool) error {
	// perform checks (light) before writing file (heavy)
	if overwrite {
		if err := os.Remove(filename); err != nil {
			log.Printf("could not remove %s due to '%s'", filename, err)
		}
	}

	// If script exists or we can't stat it
	stat, err := os.Stat(filename)

	switch {
	case os.IsNotExist(err): // nothing to do
	case err != nil:
		return fmt.Errorf("failed to determine if target file '%q' exists: %w", filename, err)
	case stat != nil && !overwrite:
		// err is nil, file exists and cannot be overwritten
		return fmt.Errorf("destination '%q' already exists", filename)
	case stat != nil && stat.IsDir(): // ensure target is not a directory
		return fmt.Errorf("destination '%q' exists and is a directory", filename)
	}

	// objective: avoid corrupting an existing file
	//
	// option 1: write to temporary file, rename and cleanup
	// - generate temporary file name for the destination
	// - rename existing file if any to a backup file
	// - rename temp file to destination file
	// - remove backup file.
	// with this option, if the destination file cannot be written, i.e., renamed
	// then the operation may fail late
	//
	// option 2 (implemented): rename existing file, write to destination, attempt recovery
	// if there's an error. This approach has a fail fast property, if the
	// destination file cannot be written, it fails right away.
	// - rename existing file if any to a backup file
	// - write to target file
	// - recover or remove backup file

	var tempName string

	if stat != nil {
		suffix, err := generateString()
		if err != nil {
			return err
		}

		tempName = filename + suffix

		if err := os.Rename(filename, tempName); err != nil {
			return fmt.Errorf("cannot backup destination file to '%q': %w", tempName, err)
		}
	}

	if err := os.WriteFile(filename, content, 0o755); err != nil {
		// cleanup
		if err := os.Remove(filename); err != nil {
			log.Printf("cleanup failed: could not remove created file '%q': %s", filename, err)
		}

		if tempName != "" { // attempt recovery when the original was preserved
			if err := os.Rename(tempName, filename); err != nil {
				log.Printf("cannot backup destination file to '%q': %s", tempName, err)
			}
		}

		return fmt.Errorf("cannot create destination file '%q': %w", filename, err)
	}

	if tempName == "" {
		return nil
	}

	// cleanup
	if err := os.Remove(tempName); err != nil {
		return fmt.Errorf("cleanup failed, could not remove temporary file '%q': %w", tempName, err)
	}

	return nil
}

// generates an 8-character long string using a base32 lowercase alphabet.
func generateString() (string, error) {
	var bytes [5]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return "", fmt.Errorf("cannot generate random string: %w", err)
	}

	return strings.ToLower(base32.StdEncoding.EncodeToString(bytes[:])), nil
}
