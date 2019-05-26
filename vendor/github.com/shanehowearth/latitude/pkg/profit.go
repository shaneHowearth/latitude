package latitude

import "fmt"

// OFFSET - Array index offset
const OFFSET = 1

// GetMaxProfit -
func GetMaxProfit(vals []int) (minIdx, maxIdx, profit int, err error) {
	if len(vals) < 2 {
		return -1, -1, 0, fmt.Errorf("not enough values to determine a profit")
	}
	minIdx, maxIdx, profit, err = GetLocalProfit(0, 0, vals)
	if err != nil {
		return -1, -1, 0, err
	}
	for idx := maxIdx + OFFSET; idx < len(vals)-1; {
		localMinIdx, localMaxIdx, localProfit, err := GetLocalProfit(idx, minIdx, vals)
		if err != nil {
			break
		}
		if localProfit > profit {
			minIdx, maxIdx, profit = localMinIdx, localMaxIdx, localProfit
		}
		// Update i here because we want to stop this loop
		// at the correct index in the vals slice.
		// But it is not a uniform increment.
		idx += localMaxIdx + OFFSET
	}
	return minIdx, maxIdx, profit, nil
}

// GetLocalProfit - max profit available in the given slice
func GetLocalProfit(idx, currentMinIdx int, vals []int) (minIdx, maxIdx, profit int, err error) {
	if len(vals) < 2 {
		return -1, -1, 0, fmt.Errorf("not enough values supplied to find the local profit")
	}
	maxIdx, err = FindMax(vals[idx+OFFSET:])
	if err != nil {
		return -1, -1, 0, err
	}
	minIdx, err = FindMin(vals[:maxIdx+idx+OFFSET])
	if err != nil {
		return -1, -1, 0, err
	}
	if vals[currentMinIdx] < vals[minIdx] {
		minIdx = currentMinIdx
	}
	profit = vals[maxIdx+idx+OFFSET] - vals[minIdx]
	return minIdx, maxIdx + idx + OFFSET, profit, nil
}

// FindMax - Find highest value in provided slice, return the index
func FindMax(vals []int) (int, error) {
	if len(vals) < 2 {
		return -1, fmt.Errorf("not enough values supplied to find the max")
	}
	idx := 0
	maxVal := vals[0]
	for i := range vals {
		if vals[i] > maxVal {
			maxVal = vals[i]
			idx = i
		}
	}
	return idx, nil
}

// FindMin - Find lowest value in provided slice, return the index
func FindMin(vals []int) (int, error) {
	if len(vals) < 1 {
		return -1, fmt.Errorf("no values supplied to find the min")
	}
	idx := 0
	minVal := vals[0]
	for i := range vals {
		if vals[i] < minVal {
			minVal = vals[i]
			idx = i
		}
	}
	return idx, nil
}
