package latitude

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
