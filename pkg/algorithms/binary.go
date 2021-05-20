package algorithms

import "strconv"

func GetAsBinaryImperative(num int) string {
	output := ""

	for i := num; i > 0; i = i / 2 {
		output = strconv.Itoa(i%2) + output
	}

	return output
}

func GetAsBinaryFunctional(num int, output string) string {
	if num == 0 {
		return output
	}

	return GetAsBinaryFunctional(num/2, strconv.Itoa(num%2)+output)
}
