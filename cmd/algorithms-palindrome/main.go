package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []string{
		"abba",
		"alice",
		"kaiak",
		"Finnif",
		"reviveR",
		"SAIPPUAKIVIKAUPPIAS",
		"revive",
		"rr",
		"ra",
		"r",
		"",
	}

	for _, input := range inputs {
		outputImperative := algorithms.IsPalindromeImperative(input)
		outputFunctional := algorithms.IsPalindromeFunctional(input, 1)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
