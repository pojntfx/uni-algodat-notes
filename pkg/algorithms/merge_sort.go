package algorithms

func MergeSortImperative(items []int) []int {
	return []int{}
}

func MergeSortFunctional(input []int) []int {
	// Array is fully split
	if len(input) < 2 {
		return input
	}

	// Split the array
	center := len(input) / 2

	left := MergeSortFunctional(input[center:])
	right := MergeSortFunctional(input[:center])

	// Merge
	output := []int{}

	// Merge left and right
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			output = append(output, left[i])

			i++
		} else {
			output = append(output, right[j])

			j++
		}
	}

	// Add remaining elements
	output = append(output, left[i:]...)
	output = append(output, right[j:]...)

	return output
}
