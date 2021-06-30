package algorithms

func QuickSortImperative(items []int) []int {
	return []int{}
}

func QuickSortFunctional(input []int) []int {
	// We have only the pivot element left
	if len(input) < 1 {
		return input
	}

	// Make the last element the pivot element
	pivot := input[len(input)-1]

	// Pop the pivot element
	input = input[:len(input)-1]

	// Split the array into sections smaller and greater than the pivot element
	smaller := []int{}
	greater := []int{}
	for _, el := range input {
		if el < pivot {
			smaller = append(smaller, el)
		} else {
			greater = append(greater, el)
		}
	}

	return append(QuickSortFunctional(smaller), append([]int{pivot}, QuickSortFunctional(greater)...)...)
}
