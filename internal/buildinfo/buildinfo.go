package buildinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime/debug"
	"slices"
	"strings"
	"sync"

	"golang.org/x/mod/semver"
)

// Version of the executable set at link-time via
// -ldflags '-X github.com/julio-lopez/goexp/internal/buildinfo.Version=foo'.
// Build information will be retrieved from the runtime when Version is not
// explicitly set at build (link) time.

var Version, Info, Repo string

// runtime.debug.ReadBuildInfo()
func PrintBuildInfo() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("could not read build info")
	}

	// fmt.Printf("buildInfo: %s\n\n", toJSONIndent(buildInfo))
	// fmt.Println("runtime.debug.ReadBuildInfo():")
	// fmt.Println(bi)

	fmt.Println("----")
	fmt.Println(" GoVersion:", bi.GoVersion)
	fmt.Println(" Path:", bi.Path)
	fmt.Println(" Main.Path:", bi.Main.Path)
	fmt.Println(" Main.Version:", bi.Main.Version)
	fmt.Println(" Main.Sum:", bi.Main.Sum)

	fmt.Println(" BuildSetting(vcs):", getBuildSettingValue(bi, "vcs"))
	vr, vt, modified := getRevision(bi.Settings)
	fmt.Println(" BuildSetting(vcs.revision):", vr)
	fmt.Println(" BuildSetting(vcs.time):", vt)
	fmt.Println(" BuildSetting(vcs.modified):", modified)
	fmt.Println("----")

	fmt.Println("Version:", getVersion())
	fmt.Println("Version:", getVersion())

	fmt.Println("build version:", semver.Build(bi.Main.Version))
	fmt.Println("prerelease version:", semver.Prerelease(bi.Main.Version))
	fmt.Println("computed build info:", getRevisionString(bi.Settings))

	// fmt.Printf("%q\n", strings.Split(bi.Main.Version, "."))
}

func init() {
	log.Println("init() Version:", Version)
	log.Println("init() Info:", Info)
	log.Println("init() Repo:", Repo)

	v := initVersion()

	Version = v
}

func initVersion() string {
	if Version != "" {
		return Version
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}

	return "(unknown version)"
}

var getVersionOnce = sync.OnceValue(initVersion)

func getVersion() string {
	return getVersionOnce()
}

func getBuildSettingValue(b *debug.BuildInfo, k string) string {
	v, _ := getBuildSetting(b, k)

	return v
}

func getBuildSetting(b *debug.BuildInfo, k string) (string, bool) {
	i := slices.IndexFunc(b.Settings, func(setting debug.BuildSetting) bool {
		return setting.Key == k
	})

	if i >= 0 {
		return b.Settings[i].Value, true
	}

	return "", false
}

func getRevision(s []debug.BuildSetting) (revision, vcsTime string, vcsModified bool) {
	for _, v := range s {
		switch v.Key {
		case "vcs.revision":
			revision = v.Value
		case "vcs.time":
			vcsTime = v.Value
		case "vcs.modified":
			if strings.ToLower(v.Value) == "true" {
				vcsModified = true
			}
		}
	}

	return
}

func getRevisionString(s []debug.BuildSetting) string {
	r, t, modified := getRevision(s)

	const maxRevLength = 12

	switch {
	case r == "":
		r = "(unknown revision)"
	case len(r) > maxRevLength:
		r = r[0:maxRevLength]
	}

	var modStr string

	if modified {
		modStr = "+dirty"
	}

	return "0." + t + "-" + r + modStr
}

func toJSONIndent(v any) []byte {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatalln("could not JSON-encode build info:", err)
	}

	return b
}
