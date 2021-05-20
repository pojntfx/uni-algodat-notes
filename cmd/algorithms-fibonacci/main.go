package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []int{5, 6, 9, 15, 31}

	for _, input := range inputs {
		outputImperative := algorithms.GetFibonacciImperative(input)
		outputFunctional := algorithms.GetFibonacciFunctional(input)

		fmt.Printf("in=%-3v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
