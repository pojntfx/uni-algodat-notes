package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []int{5, 6, 9, 15, 31, 66, 242}

	for _, input := range inputs {
		outputImperative := algorithms.GetAsBinaryImperative(input)
		outputFunctional := algorithms.GetAsBinaryFunctional(input, "")

		fmt.Printf("dec=%-3v bin=%08v=%08v\n", input, outputImperative, outputFunctional)
	}
}
