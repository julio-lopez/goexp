package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/julio-lopez/goexp/internal/junk/scripts"
)

func main() {
	base := os.Getenv("HOME")
	// base := os.TempDir()
	name := filepath.Join(base, "tmp", "scripts", "file")

	content := []string{
		"echo 'hello!' ;",
		"echo foo | cat",
		"\n# comment",
		"\nexit",
	}

	if err := scripts.WriteScript(name, content, false); err != nil {
		log.Fatalln("error:", err)
	}
}
