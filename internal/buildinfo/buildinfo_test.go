package buildinfo

import (
	"encoding/json"
	"log"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRevisionString(t *testing.T) {
	cases := []struct {
		input []debug.BuildSetting
		want  string
	}{
		{
			want: "0.-(unknown revision)",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.modified",
					Value: "true",
				},
			},
			want: "0.-(unknown revision)+dirty",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
			},
			want: "0.2025-04-12T16:01:30Z-(unknown revision)",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
				{
					Key:   "vcs.modified",
					Value: "true",
				},
			},
			want: "0.2025-04-12T16:01:30Z-(unknown revision)+dirty",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
				{
					Key:   "vcs.revision",
					Value: "353676da445938316fa00b2b812a61f4b1dd3ffa",
				},
			},
			want: "0.2025-04-12T16:01:30Z-353676da4459",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
				{
					Key:   "vcs.revision",
					Value: "353676da4459",
				},
			},
			want: "0.2025-04-12T16:01:30Z-353676da4459",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
				{
					Key:   "vcs.revision",
					Value: "353676da",
				},
			},
			want: "0.2025-04-12T16:01:30Z-353676da",
		},
		{
			input: []debug.BuildSetting{
				{
					Key:   "vcs.time",
					Value: "2025-04-12T16:01:30Z",
				},
				{
					Key:   "vcs.revision",
					Value: "353676da445938316fa00b2b812a61f4b1dd3ffa",
				},
				{
					Key:   "vcs.modified",
					Value: "true",
				},
			},
			want: "0.2025-04-12T16:01:30Z-353676da4459+dirty",
		},
	}

	for _, c := range cases {
		got := getRevisionString(c.input)
		require.Equal(t, c.want, got)
	}
}

func TestReadBuildInfo(t *testing.T) {
	bi, ok := debug.ReadBuildInfo()
	require.True(t, ok)
	require.NotNil(t, bi)

	t.Logf("%s", toJSONIndent(bi))
}

func toJSONIndent(v any) []byte {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatalln("could not JSON-encode build info:", err)
	}

	return b
}
