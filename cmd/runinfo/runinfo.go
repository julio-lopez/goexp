package main

import (
	"fmt"
	"runtime/debug"
	"sync"
)

// version of the executable set at link-time. If not set, then the info will
// be retrieved from the runtime
var version string

func main() {
	fmt.Println("version 1:", getVersion2())
	fmt.Println("version 2:", getVersion2())
}

var getVersion = sync.OnceValue(func() string {
	if version != "" {
		return version
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}

	return "(unknown version)"
})

func getVersion2() string {
	return getVersion()
}
