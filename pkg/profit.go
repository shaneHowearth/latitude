package latitude

// OFFSET - Array index offset
const OFFSET = 1

// GetMaxProfit -
func GetMaxProfit(vals []int) (minIdx, maxIdx, profit int) {
	if len(vals) < 2 {
		return -1, -1, 0
	}
	minIdx, maxIdx, profit = GetLocalProfit(0, 0, vals)
	if minIdx == -1 || maxIdx == -1 {
		return -1, -1, 0
	}
	for idx := maxIdx + OFFSET; idx < len(vals)-1; {
		localMinIdx, localMaxIdx, localProfit := GetLocalProfit(idx, minIdx, vals)
		if localProfit > profit {
			minIdx, maxIdx, profit = localMinIdx, localMaxIdx, localProfit
		}
		if localMaxIdx == -1 {
			break
		}
		// Update i here because we want to stop this loop
		// at the correct index in the vals slice.
		// But it is not a uniform increment.
		idx += localMaxIdx + OFFSET
	}
	return minIdx, maxIdx, profit
}

// GetLocalProfit - max profit available in the given slice
func GetLocalProfit(idx, currentMinIdx int, vals []int) (minIdx, maxIdx, profit int) {
	if len(vals) < 2 {

		return -1, -1, 0
	}
	maxIdx = FindMax(vals[idx+OFFSET:])
	if maxIdx != -1 {
		minIdx = FindMin(vals[:maxIdx+idx+OFFSET])
	}
	if maxIdx == -1 || minIdx == -1 {
		return -1, -1, 0
	}
	if vals[currentMinIdx] < vals[minIdx] {
		minIdx = currentMinIdx
	}
	profit = vals[maxIdx+idx+OFFSET] - vals[minIdx]
	return minIdx, maxIdx + idx + OFFSET, profit
}

// FindMax - Find highest value in provided slice, return the index
func FindMax(vals []int) int {
	if len(vals) < 2 {
		return -1
	}
	idx := 0
	maxVal := vals[0]
	for i := range vals {
		if vals[i] > maxVal {
			maxVal = vals[i]
			idx = i
		}
	}
	return idx
}

// FindMin - Find lowest value in provided slice, return the index
func FindMin(vals []int) int {
	if len(vals) < 1 {
		return -1
	}
	idx := 0
	minVal := vals[0]
	for i := range vals {
		if vals[i] < minVal {
			minVal = vals[i]
			idx = i
		}
	}
	return idx
}
