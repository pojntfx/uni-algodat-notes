package algorithms

func SequentialSearchImperative(haystack []int, needle int) int {
	for i, candidate := range haystack {
		if candidate == needle {
			return i
		}
	}

	return -1
}

func SequentialSearchFunctional(haystack []int, needle int, i int) int {
	if i == len(haystack)-1 {
		return -1
	}

	if haystack[i] == needle {
		return i
	}

	return SequentialSearchFunctional(haystack, needle, i+1)
}
