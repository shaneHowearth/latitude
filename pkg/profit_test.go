package latitude_test

import (
	latitude "latitude/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMax(t *testing.T) {
	testcases := map[string]struct {
		inputVals []int
		output    int
	}{
		"Single Max": {
			inputVals: []int{1, 9},
			output:    1,
		},
		"Two Max Vals": {
			inputVals: []int{1, 9, 9},
			output:    1,
		},
		"First Max Val": {
			inputVals: []int{10, 9, 9},
			output:    0,
		},
		"Negative Max Val": {
			inputVals: []int{-10, -9, -9},
			output:    1,
		},
		"Empty Slice": {
			inputVals: []int{},
			output:    -1,
		},
	}

	for name, tc := range testcases {
		out := latitude.FindMax(tc.inputVals)
		assert.Equal(t, tc.output, out, "Output produced incorrect val for %s", name)
	}
}

func TestFindMin(t *testing.T) {
	testcases := map[string]struct {
		inputVals []int
		output    int
	}{
		"Single Min": {
			inputVals: []int{1, 9},
			output:    0,
		},
		"Two Min Vals": {
			inputVals: []int{10, 9, 9},
			output:    1,
		},
		"First Max Val": {
			inputVals: []int{1, 9, 9},
			output:    0,
		},
		"Negative Min Val": {
			inputVals: []int{-10, -9, -9},
			output:    0,
		},
		"Empty Slice": {
			inputVals: []int{},
			output:    -1,
		},
	}

	for name, tc := range testcases {
		out := latitude.FindMin(tc.inputVals)
		assert.Equal(t, tc.output, out, "Output produced incorrect val for %s", name)
	}
}
