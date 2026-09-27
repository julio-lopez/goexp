package intrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetKeyRange(t *testing.T) {
	cases := []struct {
		input   map[int]bool
		want    closedIntRange
		length  uint
		isEmpty bool
	}{
		{
			isEmpty: true,
			want:    closedIntRange{lo: 0, hi: -1},
		},
		{
			input:  map[int]bool{0: true},
			want:   closedIntRange{lo: 0, hi: 0},
			length: 1,
		},
		{
			input:  map[int]bool{-5: true},
			want:   closedIntRange{lo: -5, hi: -5},
			length: 1,
		},
		{
			input:  map[int]bool{-5: true, -4: true},
			want:   closedIntRange{lo: -5, hi: -4},
			length: 2,
		},
		{
			input:  map[int]bool{0: true},
			want:   closedIntRange{lo: 0, hi: 0},
			length: 1,
		},
		{
			input:  map[int]bool{5: true},
			want:   closedIntRange{lo: 5, hi: 5},
			length: 1,
		},
		{
			input:  map[int]bool{0: true, 1: true},
			want:   closedIntRange{lo: 0, hi: 1},
			length: 2,
		},
		{
			input:  map[int]bool{8: true, 9: true},
			want:   closedIntRange{lo: 8, hi: 9},
			length: 2,
		},
		{
			input:  map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true},
			want:   closedIntRange{lo: 1, hi: 5},
			length: 5,
		},
		{
			input:  map[int]bool{8: true, 10: true},
			want:   closedIntRange{lo: 8, hi: 10},
			length: 3,
		},
		{
			input:  map[int]bool{1: true, 2: true, 3: true, 5: true},
			want:   closedIntRange{lo: 1, hi: 5},
			length: 5,
		},
		{
			input:  map[int]bool{-5: true, -7: true},
			want:   closedIntRange{lo: -7, hi: -5},
			length: 3,
		},
		{
			input:  map[int]bool{0: true, minInt: true},
			want:   closedIntRange{lo: minInt, hi: 0},
			length: -minInt + 1,
		},
		{
			input:  map[int]bool{0: true, maxInt: true},
			want:   closedIntRange{lo: 0, hi: maxInt},
			length: maxInt + 1,
		},
		{
			input:   map[int]bool{maxInt: true, minInt: true},
			want:    closedIntRange{lo: minInt, hi: maxInt},
			length:  0, // corner case, not representable :(
			isEmpty: true,
		},
		{
			input:  map[int]bool{minInt: true},
			want:   closedIntRange{lo: minInt, hi: minInt},
			length: 1,
		},
		{
			input:  map[int]bool{maxInt - 1: true},
			want:   closedIntRange{lo: maxInt - 1, hi: maxInt - 1},
			length: 1,
		},
		{
			input:  map[int]bool{maxInt: true},
			want:   closedIntRange{lo: maxInt, hi: maxInt},
			length: 1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint("case: ", i), func(t *testing.T) {
			got := getKeyRange(tc.input)

			require.Equal(t, tc.want, got, "input: %#v", tc.input)
			require.Equal(t, tc.length, got.length())
			require.Equal(t, tc.isEmpty, got.isEmpty())
		})
	}
}

func TestGetContiguosKeyRange(t *testing.T) {
	invalidEmptyRange := closedIntRange{-1, -2}

	cases := []struct {
		input     map[int]bool
		want      closedIntRange
		shouldErr bool
		length    uint
		isEmpty   bool
	}{
		{
			isEmpty: true,
			want:    closedIntRange{0, -1},
		},
		{
			input:  map[int]bool{0: true},
			want:   closedIntRange{lo: 0, hi: 0},
			length: 1,
		},
		{
			input:  map[int]bool{-5: true},
			want:   closedIntRange{lo: -5, hi: -5},
			length: 1,
		},
		{
			input:  map[int]bool{-5: true, -4: true},
			want:   closedIntRange{lo: -5, hi: -4},
			length: 2,
		},
		{
			input:  map[int]bool{0: true},
			want:   closedIntRange{lo: 0, hi: 0},
			length: 1,
		},
		{
			input:  map[int]bool{5: true},
			want:   closedIntRange{lo: 5, hi: 5},
			length: 1,
		},
		{
			input:  map[int]bool{0: true, 1: true},
			want:   closedIntRange{lo: 0, hi: 1},
			length: 2,
		},
		{
			input:  map[int]bool{8: true, 9: true},
			want:   closedIntRange{lo: 8, hi: 9},
			length: 2,
		},
		{
			input:  map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true},
			want:   closedIntRange{lo: 1, hi: 5},
			length: 5,
		},
		{
			input:     map[int]bool{8: true, 10: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:     map[int]bool{1: true, 2: true, 3: true, 5: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:     map[int]bool{-5: true, -7: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:     map[int]bool{0: true, minInt: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:     map[int]bool{0: true, maxInt: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:     map[int]bool{maxInt: true, minInt: true},
			want:      invalidEmptyRange,
			shouldErr: true,
			isEmpty:   true,
		},
		{
			input:  map[int]bool{minInt: true},
			want:   closedIntRange{lo: minInt, hi: minInt},
			length: 1,
		},
		{
			input:  map[int]bool{maxInt - 1: true},
			want:   closedIntRange{lo: maxInt - 1, hi: maxInt - 1},
			length: 1,
		},
		{
			input:  map[int]bool{maxInt: true},
			want:   closedIntRange{lo: maxInt, hi: maxInt},
			length: 1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint("case: ", i), func(t *testing.T) {
			got, err := getContiguousKeyRange(tc.input)
			if tc.shouldErr {
				require.Error(t, err, "input: %v", tc.input)
			}

			require.Equal(t, tc.want, got, "input: %#v", tc.input)
			require.Equal(t, tc.length, got.length())
			require.Equal(t, tc.isEmpty, got.isEmpty())
		})
	}
}
