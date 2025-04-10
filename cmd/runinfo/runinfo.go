package main

import (
	"github.com/julio-lopez/goexp/internal/buildinfo"
)

// runtime.debug.ReadBuildInfo()
func main() {
	buildinfo.PrintBuildInfo()
}
