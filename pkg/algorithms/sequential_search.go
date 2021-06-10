package algorithms

func SequentialSearchImperative(input []int, target int) int {
	for i, candidate := range input {
		if candidate == target {
			return i
		}
	}

	return -1
}

func SequentialSearchFunctional(input []int, target int, i int) int {
	if i == len(input)-1 {
		return -1
	}

	if input[i] == target {
		return i
	}

	return SequentialSearchFunctional(input, target, i+1)
}
