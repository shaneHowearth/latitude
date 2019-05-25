package latitude

// LocalProfit - max profit available in the given slice
func LocalProfit(vals []int) (min, max, profit int) {
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
