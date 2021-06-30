package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{
		{1, 6, 10, 2, 11},
		{2, 1, 2, 9, 3, 5},
		{3, 38, 4, 45, 5, 7, 12, 50},
	}

	for _, input := range inputs {
		outputImperative := algorithms.MergeSortImperative(input)
		outputFunctional := algorithms.MergeSortFunctional(input)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
