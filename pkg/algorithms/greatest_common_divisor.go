package algorithms

func GetGreatestCommonDivisorImperative(a int, b int) int {
	for b != 0 {
		if a > b {
			a, b = b, a%b
		} else {
			b = b % a
		}
	}

	return a
}

func GetGreatestCommonDivisorFunctional(a int, b int) int {
	if b == 0 {
		return a
	}

	if a > b {
		return GetGreatestCommonDivisorFunctional(b, a%b)
	}

	return GetGreatestCommonDivisorFunctional(a, b%a)
}
