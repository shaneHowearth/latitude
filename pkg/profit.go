package latitude

// OFFSET - Array index offset
const OFFSET = 1

// GetMaxProfit -
func GetMaxProfit(vals []int) (min, max, profit int) {
	min, max, profit = GetLocalProfit(vals)
	if min == -1 || max == -1 {
		return -1, -1, 0
	}
	for i := max + OFFSET; i < len(vals)-1; {
		localMin, localMax, localProfit := GetLocalProfit(vals[i:])
		if localProfit > profit {
			min, max, profit = max+localMin+OFFSET, max+localMax+OFFSET, localProfit
		}
		if localMax == -1 {
			break
		}
		// Update i here because we want to stop this loop
		// at the correct index in the vals slice.
		// But it is not a uniform increment.
		i += localMax
	}
	return min, max, profit
}

// GetLocalProfit - max profit available in the given slice
func GetLocalProfit(vals []int) (min, max, profit int) {
	max = FindMax(vals)
	if max != -1 {
		min = FindMin(vals[0:max])
	}
	if max == -1 || min == -1 {
		return -1, -1, 0
	}
	profit = vals[max] - vals[min]
	return min, max, profit
}

// FindMax - Find highest value in provided slice, return the index
func FindMax(vals []int) int {
	if len(vals) < 1 {
		return -1
	}
	idx := 0
	max := vals[0]
	for i := range vals {
		if vals[i] > max {
			max = vals[i]
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
	min := vals[0]
	for i := range vals {
		if vals[i] < min {
			min = vals[i]
			idx = i
		}
	}
	return idx
}
