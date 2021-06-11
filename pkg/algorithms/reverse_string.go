package algorithms

import (
	"fmt"
	"strings"
)

func GetReversedStringImperative(input string) string {
	output := ""

	for _, character := range strings.Split(input, "") {
		output = character + output
	}

	return output
}

func GetReversedStringFunctional(input string, output string, i int) string {
	if i == len(input) {
		return output
	}

	return GetReversedStringFunctional(input, fmt.Sprintf("%c", input[i])+output, i+1)
}
