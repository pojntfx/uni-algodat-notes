package algorithms

func GetSumImperative(max int) int {
	sum := 0

	for i := 1; i <= max; i++ {
		sum += i
	}

	return sum
}

func GetSumFunctional(max int, sum int) int {
	if max == 0 {
		return sum
	}

	return GetSumFunctional(max-1, sum+max)
}

// Alternative implementation, based on conversion from GetSumImperative
// func GetSumFunctional(max int, sum int, i int) int {
// 	if i == max {
// 		return sum
// 	}

// 	return GetSumFunctional(max, sum+i, i+1)
// }
