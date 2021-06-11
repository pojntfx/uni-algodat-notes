package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []string{
		"alice",
		"bob",
		"teddy",
		"alan",
	}

	for _, input := range inputs {
		outputImperative := algorithms.GetReversedStringImperative(input)
		outputFunctional := algorithms.GetReversedStringFunctional(input, "", 0)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
