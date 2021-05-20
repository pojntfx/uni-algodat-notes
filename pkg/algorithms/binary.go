package algorithms

import "strconv"

func GetAsBinaryImperative(num int) string {
	output := ""

	for i := num; i > 0; i = i / 2 {
		output = strconv.Itoa(i%2) + output
	}

	return output
}
