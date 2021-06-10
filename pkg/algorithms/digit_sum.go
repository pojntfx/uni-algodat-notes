package algorithms

func GetDigitSumImperative(num int) int {
	sum := 0

	curr := num
	for curr > 0 {
		sum += curr % 10

		curr = curr / 10
	}

	return sum
}

func GetDigitSumFunctional(num int, sum int) int {
	if num == 0 {
		return sum
	}

	return GetDigitSumFunctional(num/10, sum+(num%10))
}
