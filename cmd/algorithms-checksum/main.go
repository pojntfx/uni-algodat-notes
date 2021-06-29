package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{
		{1, 3, 6, 6},
		{3, 4},
		{10, 15},
		{12, 17},
		{9, 18},
		{156, 66},
	}

	for _, input := range inputs {
		outputImperative := algorithms.GetChecksumImperative(input)
		outputFunctional := algorithms.GetChecksumFunctional(input, 0, 0)

		fmt.Printf("in=%v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
