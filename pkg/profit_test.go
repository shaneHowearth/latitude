package latitude_test

import (
	latitude "latitude/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxProfit(t *testing.T) {
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
			min:       1,
			max:       4,
			profit:    3,
		},
		"First Max": {
			inputVals: []int{4, 5, 3, 4, 5},
			min:       2,
			max:       4,
			profit:    2,
		},
		"Multi max check": {
			inputVals: []int{4, 7, 3, 4, 6, 1, 2, 4, 3, 1, 5, 4},
			min:       5,
			max:       10,
			profit:    4,
		},
		"Long Input": {
			inputVals: []int{54, 40, 47, 88, 37, 81, 32, 67, 56, 93,
				18, 16, 12, 95, 96, 17, 11, 4, 16, 66,
				51, 18, 6, 56, 16, 31, 38, 49, 6, 15,
				80, 65, 57, 70, 43, 7, 21, 12, 62, 72,
				53, 51, 94, 90, 4, 73, 29, 14, 72, 32,
				72, 63, 85, 48, 49, 18, 11, 60, 10, 89,
				16, 75, 16, 61, 51, 28, 83, 76, 42, 3,
				37, 53, 96, 49, 95, 1, 94, 60, 94, 66, 89,
				4, 27, 70, 25, 53, 14, 83, 14, 85, 26,
				25, 92, 22, 56, 47, 38, 73, 95, 71, 49},
			min:    75,
			max:    98,
			profit: 94,
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
		minOut, maxOut, profitOut := latitude.GetMaxProfit(tc.inputVals)
		assert.Equal(t, tc.min, minOut, "Output produced incorrect min index for %s", name)
		assert.Equal(t, tc.max, maxOut, "Output produced incorrect max index for %s", name)
		assert.Equal(t, tc.profit, profitOut, "Output produced incorrect profit for %s", name)
	}

}

func TestGetLocalProfit(t *testing.T) {
	testcases := map[string]struct {
		inputVals []int
		min       int
		max       int
		profit    int
		err       error
	}{
		"Happy Path": {
			inputVals: []int{1, 2, 3, 4, 5},
			min:       0,
			max:       4,
			profit:    4,
		},
		"No result": {
			inputVals: []int{9, 9},
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
		"Single entry list": {
			inputVals: []int{5},
			min:       -1,
			max:       -1,
			profit:    0,
		},
		"Empty list": {
			inputVals: []int{},
			min:       -1,
			max:       -1,
			profit:    0,
		},
	}

	for name, tc := range testcases {
		minOut, maxOut, profitOut := latitude.GetLocalProfit(0, 0, tc.inputVals)
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
