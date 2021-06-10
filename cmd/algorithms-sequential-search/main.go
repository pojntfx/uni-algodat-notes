package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{{5, 0, 6, 15, 31}, {1, 2, 3, 4, 5}}
	target := 6

	for _, input := range inputs {
		outputImperative := algorithms.SequentialSearchImperative(input, target)
		outputFunctional := algorithms.SequentialSearchFunctional(input, target, 0)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
