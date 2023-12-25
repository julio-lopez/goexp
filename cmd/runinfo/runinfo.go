package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

// Version of the executable set at link-time. If not set, then the info will
// be retrieved from the runtime
var Version string

func main() {
	fmt.Println("version 1:", getVersion())
	fmt.Println("version 2:", getVersion())
}

var getVersion func() string = sync.OnceValue(func() string {
	log.Println("getVersion called")
	if Version != "" {
		return Version
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}

	return "(unknown version)"
})
