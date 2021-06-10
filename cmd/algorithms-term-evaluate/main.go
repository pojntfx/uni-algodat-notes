package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []string{
		"5+5",
		"5-5",
		"10-5+6",
		"20*5",
		"20*5/5",
	}

	for _, input := range inputs {
		outputImperative := algorithms.TermEvaluateImperative(input)
		outputFunctional := algorithms.TermEvaluateFunctional(input)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
