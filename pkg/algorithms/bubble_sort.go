package algorithms

func BubbleSortImperative(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)

	for j := 0; j < len(input); j++ {
		for i := 1; i < len(input)-j; i++ {
			if output[i] < output[i-1] {
				output[i], output[i-1] = output[i-1], output[i]
			}
		}
	}

	return output
}

// Alternative implementation with flag variable
// func BubbleSortImperative(input []int) []int {
// 	output := make([]int, len(input))
// 	copy(output, input)

// 	for {
// 		sorted := true

// 		for i := range output {
// 			if i == len(output)-1 {
// 				if sorted {
// 					return output
// 				}

// 				continue
// 			}

// 			if output[i] > output[i+1] {
// 				sorted = false

// 				output[i], output[i+1] = output[i+1], output[i]
// 			}
// 		}
// 	}
// }

func BubbleSortFunctional(output []int, j int) []int {
	if j >= len(output) {
		return output
	}

	return BubbleSortFunctional(BubbleSortFunctionalNested(output, j, 1), j+1)
}

func BubbleSortFunctionalNested(output []int, j int, i int) []int {
	if i >= len(output) {
		return output
	}

	if output[i] < output[i-1] {
		// This can't be made side-effect free due to syntax limitations
		output[i], output[i-1] = output[i-1], output[i]
	}

	return BubbleSortFunctionalNested(output, j, i+1)
}
