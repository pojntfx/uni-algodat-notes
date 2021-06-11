package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{
		{3, 4},
		{10, 15},
		{12, 17},
		{9, 18},
		{156, 66},
	}

	for _, input := range inputs {
		outputImperative := algorithms.GetGreatestCommonDivisorImperative(input[0], input[1])
		outputFunctional := algorithms.GetGreatestCommonDivisorFunctional(input[0], input[1])

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
