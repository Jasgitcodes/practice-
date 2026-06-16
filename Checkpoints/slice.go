package piscine

func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	// If no indices are given, the operation is invalid
	// so we return nil as required by the exercise
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0] // first index = start
	end := n         // default end = full length of slice

	// If a second index is provided, use it as end
	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Handle negative start index (count from end)
	if start < 0 {
		start += n
	}

	// Handle negative end index (count from end)
	if end < 0 {
		end += n
	}

	// If start goes below 0 after adjustment, clamp it to 0
	if start < 0 {
		start = 0
	}

	// If end goes beyond slice length, clamp it to n
	if end > n {
		end = n
	}

	// If range is invalid or empty, return nil
	if start >= end {
		return nil
	}

	// Return the sliced portion directly
	return a[start:end]
}
