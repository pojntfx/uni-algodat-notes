package algorithms

func BinarySearchFunctional(haystack []int, needle int, low int, high int) int {
	// Needle not in haystack
	if low > high {
		return -1
	}

	/*
		To make this fully side-effect free, replace all references
		to the following with these expressions
	*/
	focus := (low + high) / 2
	center := haystack[focus]

	// Needle at focus
	if needle == center {
		return focus
	}

	// Take left part
	if needle < center {
		return BinarySearchFunctional(haystack, needle, low, focus-1)
	}

	// Take right part
	return BinarySearchFunctional(haystack, needle, focus+1, high)
}

func BinarySearchImperative(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1

	for {
		// Needle not in haystack
		if low > high {
			return -1
		}

		focus := (low + high) / 2
		center := haystack[focus]

		// Needle at focus
		if needle == center {
			return focus
		}

		// Take left part
		if needle < center {
			high = focus - 1
		} else {
			// Take right part
			low = focus + 1
		}
	}
}
