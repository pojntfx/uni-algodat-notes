package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{
		{1, 6, 9, 10, 11},
		{1, 2, 3, 4, 5},
		{3, 4, 5, 7, 12, 50},
		{3, 4, 6, 7, 12, 50},
		{1, 2, 3, 4, 5, 6, 12, 50},
	}
	target := 6

	for _, input := range inputs {
		outputImperative := algorithms.BinarySearchImperative(input, target)
		outputFunctional := algorithms.BinarySearchFunctional(input, target, 0, len(input)-1)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
