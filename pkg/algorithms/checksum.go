package algorithms

func GetChecksumImperative(input []int) int {
	res := 0

	for _, num := range input {
		res += num
	}

	return res
}

func GetChecksumFunctional(input []int, res int, i int) int {
	if i >= len(input) {
		return res
	}

	return GetChecksumFunctional(input, res+input[i], i+1)
}
