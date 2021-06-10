package algorithms

import (
	"log"
	"strconv"
	"strings"
)

func TermEvaluateImperative(term string) int {
	parts := []string{}

	lastNum := ""
	for i, character := range strings.Split(term, "") {
		if character == "+" || character == "-" || character == "*" || character == "/" {
			parts = append(parts, lastNum)
			lastNum = ""

			parts = append(parts, character)

			continue
		}

		lastNum = lastNum + character

		if i == len(term)-1 {
			parts = append(parts, lastNum)
		}
	}

	log.Println(parts)

	res := 0
	lastInstruction := "+"
	for _, instruction := range parts {
		num, err := strconv.Atoi(instruction)
		if err != nil {
			lastInstruction = instruction

			continue
		}

		switch lastInstruction {
		case "+":
			res += num
		case "-":
			res -= num
		case "*":
			res *= num
		case "/":
			res /= num
		}
	}

	return res
}

func TermEvaluateFunctional(term string) int {
	// TODO: Add functional implementation using paradigm conversion

	return 0
}
