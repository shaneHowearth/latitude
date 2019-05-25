package latitude_test

import (
	latitude "latitude/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalProfit(t *testing.T) {
	testcases := map[string]struct {
		inputVals []int
		min       int
		max       int
		profit    int
	}{
		"Happy Path": {
			inputVals: []int{1, 2, 3, 4, 5},
			min:       0,
			max:       4,
			profit:    4,
		},
		"No result": {
			inputVals: []int{9, 2, 3, 4, 5},
			min:       -1,
			max:       -1,
			profit:    0,
		},
		// This is to show that the Local Profit is only getting the
		// first possible max profit
		"First Max": {
			inputVals: []int{4, 5, 3, 4, 5},
			min:       0,
			max:       1,
			profit:    1,
		},
		"Min further inside": {
			inputVals: []int{5, 2, 3, 4, 9, 3},
			min:       1,
			max:       4,
			profit:    7,
		},
		"Empty list": {
			inputVals: []int{},
			min:       -1,
			max:       -1,
			profit:    0,
		},
	}

	for name, tc := range testcases {
		minOut, maxOut, profitOut := latitude.LocalProfit(tc.inputVals)
		assert.Equal(t, tc.min, minOut, "Output produced incorrect min index for %s", name)
		assert.Equal(t, tc.max, maxOut, "Output produced incorrect max index for %s", name)
		assert.Equal(t, tc.profit, profitOut, "Output produced incorrect profit for %s", name)
	}
}

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
