package algorithms

func IsPalindromeImperative(input string) bool {
	if len(input) == 0 {
		return false
	}

	if len(input) == 1 {
		return true
	}

	for i := 1; i <= len(input); i++ {
		if input[i-1] != input[len(input)-i] {
			return false
		}
	}

	return true
}

func IsPalindromeFunctional(input string, i int) bool {
	if len(input) == 0 {
		return false
	}

	if len(input) == 1 {
		return true
	}

	if i != 1 && i <= len(input) {
		return true
	}

	if input[i-1] != input[len(input)-i] {
		return false
	}

	return IsPalindromeFunctional(input, i+1)
}
