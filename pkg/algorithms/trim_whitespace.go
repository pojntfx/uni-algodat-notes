package algorithms

import (
	"strings"
)

func TrimWhitespaceImperative(input string) string {
	output := ""

	for _, character := range strings.Split(input, "") {
		if !isWhitespace(character) {
			output += character
		}
	}

	return output
}

func TrimWhitespaceFunctional(input string, output string, i int) string {
	if i >= len(input) {
		return output
	}

	/*
		To make this fully side-effect free, replace all references
		to the following with these expressions
	*/
	curr := strings.Split(input, "")[i]

	if !isWhitespace(curr) {
		return TrimWhitespaceFunctional(input, output+curr, i+1)
	}

	return TrimWhitespaceFunctional(input, output, i+1)
}

func isWhitespace(in string) bool {
	return strings.TrimSpace(in) == ""
}
